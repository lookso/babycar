package server

import (
	babyApiV1 "babycare/api/baby/v1"
	carApiV1 "babycare/api/car/v1"
	"babycare/internal/conf"
	"babycare/internal/server/middleware"
	"babycare/internal/service/baby"
	"babycare/internal/service/car"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var errorHandle *conf.ErrorHandle

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, conf *conf.Data, errorHandler *conf.ErrorHandle,carService *car.CarService, babyService *baby.BabyService, logger log.Logger) *http.Server {
	errorHandle = errorHandler

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.AddTraceToRequest(),
			validate.Validator(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	fmt.Println("conf",conf)
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

	router := srv.Route("/")
	router.GET("v1/download_comment", carService.DownloadComment)
	return srv
}
