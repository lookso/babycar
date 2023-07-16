package api

import (
	pb "babycare/api/car/v1"
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfoReply, error) {
	s.carBiz.CreateUser(ctx, req)
	return &pb.UserInfoReply{}, nil
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	id, _ := s.carBiz.GetUser(ctx, req)
	return &pb.GetUserReply{Id: id}, nil
}
func (s *Service) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *Service) SendJson(ctx context.Context, req *pb.SendJsonRequest) (*pb.SendJsonReply, error) {
	return &pb.SendJsonReply{}, nil
}
func (s *Service) AuthToken(ctx context.Context, req *pb.AuthTokenRequest) (*pb.AuthTokenReply, error) {
	return &pb.AuthTokenReply{}, nil
}
func (s *Service) GetWechatContacts(ctx context.Context, req *pb.GetWechatContactsRequest) (*pb.GetWechatContactsReply, error) {
	return &pb.GetWechatContactsReply{}, nil
}

func (s *Service) Hello(ctx http.Context) error {
	return ctx.JSON(200, map[string]string{
		"message": "hello kratos",
	})
}
