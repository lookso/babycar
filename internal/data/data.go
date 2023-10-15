package data

import (
	"babycare/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCarData,NewBabyData, NewEngineDbRw,wire.Struct(new(DataProviderCollection), "*"))

// Data . .
type Data struct {
	EngineDb *gorm.DB
	Log      *log.Helper
}

type DataProviderCollection struct {
	EngineDb EngineDb
}

type EngineDb *gorm.DB

func NewEngineDbRw(conf *conf.Data, logger log.Logger) EngineDb {
	logDb := log.NewHelper(log.With(logger, "x_module", "data/gorm-mysql"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		//Logger:                 log.NewGorm(logger),
	})
	if err != nil {
		logDb.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		logDb.Fatalf("get contact to mysql: %v", err)
	}
	sqlDb.SetMaxIdleConns(int(conf.Database.MaxIdleConns))
	sqlDb.SetMaxOpenConns(int(conf.Database.MaxOpenConns))
	sqlDb.SetConnMaxLifetime(time.Hour)

	err = db.Use(
		dbresolver.Register(dbresolver.Config{Replicas: []gorm.Dialector{mysql.Open(conf.Database.Source)}}).
			SetConnMaxLifetime(time.Hour).
			SetMaxIdleConns(int(conf.Database.MaxIdleConns)).
			SetMaxOpenConns(int(conf.Database.MaxOpenConns)),
	)
	if err != nil {
		logDb.Fatalf("failed db use to mysql: %v", err)
	}
	logDb.Info("init mysql")
	return db
}

// NewData .
func NewData(c *conf.Data, dataProvider *DataProviderCollection, logger log.Logger) (*Data, func(), error) {

	logData := log.NewHelper(log.With(logger, "x_module", "data/resource"))

	ormEngineDb := (*gorm.DB)(dataProvider.EngineDb)
	cleanup := func() {
		if ormEngineDb != nil {
			db, _ := ormEngineDb.DB()
			_ = db.Close()
		}
		logData.Info("closing the data resources")
	}

	return &Data{
		EngineDb: ormEngineDb,
		Log:      logData,
	}, cleanup, nil
}
