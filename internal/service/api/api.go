package api

import (
	pb "babycare/api/car/v1"
	"babycare/internal/biz/car"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/structpb"
)

type Service struct {
	pb.UnimplementedCarServer
	carBiz *car.CarBiz
	log     *log.Helper
}

func NewService(carBiz *car.CarBiz, logger log.Logger) *Service {
	return &Service{
		carBiz: carBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/api")),
	}
}

func (s *Service) HealthCheck(ctx context.Context, structValue *structpb.Value) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Message: "ok",
	}, nil
}
