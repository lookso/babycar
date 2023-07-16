package car

import (
	pb "babycare/api/car/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CarBiz struct {
	carRepo ICarRepo
	log      *log.Helper
}

func NewCarBiz(userRepo ICarRepo, logger log.Logger) *CarBiz {
	return &CarBiz{carRepo: userRepo, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}

func (s *CarBiz) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfoReply, error) {
	s.carRepo.Create(ctx, req.GetNickName(), req.GetPassword(), req.GetMobile())
	return &pb.UserInfoReply{}, nil
}

func (s *CarBiz) GetUser(ctx context.Context, req *pb.GetUserRequest) (int64, error) {
	id, err := s.carRepo.GetUser(ctx, req.GetId())
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CarBiz) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *CarBiz) SendJson(ctx context.Context, req *pb.SendJsonRequest) (*pb.SendJsonReply, error) {
	return &pb.SendJsonReply{}, nil
}
func (s *CarBiz) AuthToken(ctx context.Context, req *pb.AuthTokenRequest) (*pb.AuthTokenReply, error) {
	return &pb.AuthTokenReply{}, nil
}
func (s *CarBiz) GetWechatContacts(ctx context.Context, req *pb.GetWechatContactsRequest) (*pb.GetWechatContactsReply, error) {
	return &pb.GetWechatContactsReply{}, nil
}
