package baby

import (
	pb "babycare/api/baby/v1"
	"babycare/internal/biz/baby"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BabyService struct {
	pb.UnimplementedBabyServer
	babyBiz *baby.BabyBiz
	log     *log.Helper
}

func NewBabyService(babyBiz *baby.BabyBiz, logger log.Logger) *BabyService {
	return &BabyService{
		babyBiz: babyBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/api")),
	}
}
func (s *BabyService) Example(ctx context.Context, req *pb.Null) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *BabyService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	if header, ok := transport.FromServerContext(ctx); ok {
		appId := header.RequestHeader().Get("X-App-Id")
		//s.log.Infof("appId:%s", appId)

		fmt.Println("appId", appId)

		headers := make(map[string]string)
		for _, key := range header.RequestHeader().Keys() {
			if key == "X-App-Id" {
				headers[key] = "test-BB"
			} else {
				headers[key] = header.RequestHeader().Get(key)
			}
		}
		fmt.Println("headers", headers)
	}
	resp := &pb.GetUserReply{}
	resp, err := s.babyBiz.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	resp.AccessToken = "test-BBB"
	return resp, nil
}
