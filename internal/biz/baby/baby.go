package baby

import (
	pb "babycare/api/car/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BabyBiz struct {
	babyRepo IBabyRepo
	log      *log.Helper
}

func NewBabyBiz(userRepo IBabyRepo, logger log.Logger) *BabyBiz {
	return &BabyBiz{babyRepo: userRepo, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}

func (s *BabyBiz) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfoReply, error) {
	s.babyRepo.GetUser(ctx, req.GetId())
	return &pb.UserInfoReply{}, nil
}
