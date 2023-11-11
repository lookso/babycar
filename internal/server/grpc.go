package server

import (
	carV1 "babycare/api/car/v1"
	treeV1 "babycare/api/tree/v1"
	"babycare/internal/conf"
	"babycare/internal/service/car"
	"babycare/internal/service/tree"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, carService *car.CarService, treeService *tree.TreeService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	carV1.RegisterCarServer(srv, carService)
	treeV1.RegisterTreeServer(srv, treeService)
	logger.Log(log.LevelInfo, "register car service")
	return srv
}
