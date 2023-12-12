package server

import (
	babyApiV1 "babycare/api/baby/v1"
	carApiV1 "babycare/api/car/v1"
	treeApiV1 "babycare/api/tree/v1"
	"babycare/internal/biz/biz_metrics"
	"babycare/internal/conf"
	"babycare/internal/server/middleware"
	"babycare/internal/service/baby"
	"babycare/internal/service/car"
	"babycare/internal/service/tree"
	"fmt"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var errorHandle *conf.ErrorHandle

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, conf *conf.Data, errorHandler *conf.ErrorHandle, carService *car.CarService, babyService *baby.BabyService, treeService *tree.TreeService, logger log.Logger) *http.Server {
	errorHandle = errorHandler

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.AddTraceToRequest(),
			validate.Validator(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(biz_metrics.MetricSeconds)),
				metrics.WithRequests(prom.NewCounter(biz_metrics.MetricRequests))),
		),
	}
	fmt.Println("conf", conf)
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ErrorEncoder(MyErrorEncoder))
	opts = append(opts, http.ResponseEncoder(MyResponseEncoder))
	srv := http.NewServer(opts...)
	carApiV1.RegisterCarHTTPServer(srv, carService)
	babyApiV1.RegisterBabyHTTPServer(srv, babyService)
	treeApiV1.RegisterTreeHTTPServer(srv, treeService)
	router := srv.Route("/")
	router.GET("v1/download_comment", carService.DownloadComment)
	biz_metrics.Init()
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}
