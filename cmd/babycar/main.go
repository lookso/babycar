package main

import (
	"babycare/pkg/zlog"
	"flag"
	"fmt"
	"os"
	"time"

	"babycare/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
	netHttp "net/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "babycar"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// flagconf is the config flag.
	flagconf string
	env      string
	id, _    = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&env, "env", "dev", "use env, eg: -env=dev")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	zlog.Init(Name, env, bc.Log.Filename, int(bc.Log.MaxSize), int(bc.Log.MaxBackup), int(bc.Log.MaxAge), bc.Log.Compress)
	defer zlog.Sync()
	logger := log.With(zlog.NewZapLogger(zlog.STDInstance()),
		"ts", log.Timestamp(time.RFC3339Nano),
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		// 可以添加额外k,v,满足基本日志需求
	)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Error, logger)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	defer cleanup()

	go func() {
		_ = netHttp.ListenAndServe(":6060", nil)
		fmt.Println("pprof start")
	}()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
