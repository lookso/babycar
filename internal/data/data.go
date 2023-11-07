package data

import (
	"babycare/internal/conf"
	"context"
	"time"

	// "github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"babycare/pkg/zlog"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	//redisotel "github.com/redis/go-redis/extra/redisotel/v9"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/plugin/dbresolver"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCarData, NewBabyData)

// Data . .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
}

func NewDb(conf *conf.Data, logger log.Logger) (db *gorm.DB, err error) {
	logDb := log.NewHelper(log.With(logger, "x_module", "data/gorm-mysql"))
	db, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 zlog.NewGorm(logger),
	})
	if err != nil {
		logDb.Fatalf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	if conf.Database.GetDebug() {
		db = db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		logDb.Fatalf("get contact to mysql: %v", err)
		return nil, err
	}
	sqlDb.SetMaxIdleConns(int(conf.Database.MaxIdleConns))
	sqlDb.SetMaxOpenConns(int(conf.Database.MaxOpenConns))
	sqlDb.SetConnMaxLifetime(time.Hour)

	// if err := db.Use(otelgorm.NewPlugin()); err != nil {
	// 	return nil, errors.Wrap(err, "data: db.Use error")
	// }
	// err = db.Use(
	// 	// 你的应用程序不需要读写分离或负载均衡可以去掉
	// 	dbresolver.Register(dbresolver.Config{Replicas: []gorm.Dialector{mysql.Open(conf.Database.Source)}}).
	// 		SetConnMaxLifetime(time.Hour).
	// 		SetMaxIdleConns(int(conf.Database.MaxIdleConns)).
	// 		SetMaxOpenConns(int(conf.Database.MaxOpenConns)),
	// )
	if err != nil {
		logDb.Fatalf("failed db use to mysql: %v", err)
		return nil, err
	}
	logDb.Info("init mysql")
	return db, nil
}

func NewRedis(conf *conf.Data, logger log.Logger) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		log.NewHelper(logger).Info("failed opening connection to redis: %v", err)
		return nil, err
	}
	//if err := redisotel.InstrumentTracing(rdb); err != nil {
	//	return nil, errors.Wrap(err, "data: redisotel.InstrumentTracing error")
	//}
	//if err := redisotel.InstrumentMetrics(rdb); err != nil {
	//	return nil, errors.Wrap(err, "data: redisotel.InstrumentMetrics error")
	//}
	return rdb, nil
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	logData := log.NewHelper(log.With(logger, "x_module", "data/resource"))
	db, err := NewDb(c, logger)
	if err != nil {
		return nil, nil, err
	}
	rdb, err := NewRedis(c, logger)
	if err != nil {
		return nil, nil, err
	}
	d := &Data{
		db:  db,
		rdb: rdb,
		log: logData,
	}
	cleanup := func() {
		_db, err := d.db.DB()
		if err != nil {
			log.NewHelper(logger).Errorf("database close err:%+v", err)
		}
		_ = _db.Close()
		log.NewHelper(logger).Info("closing the mysql")
		_ = d.rdb.Close()
		log.NewHelper(logger).Info("closing the redis")
		logData.Info("closing the data resources success")
	}
	return d, cleanup, nil
}
