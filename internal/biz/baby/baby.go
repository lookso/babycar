package baby

import (
	pb "babycare/api/baby/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/types/known/structpb"
)

type BabyBiz struct {
	babyRepo IBabyRepo
	log      *log.Helper
}

func NewBabyBiz(userRepo IBabyRepo, logger log.Logger) *BabyBiz {
	return &BabyBiz{babyRepo: userRepo, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}

func (s *BabyBiz) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	// 创建一个 Struct 对象
	resp := &pb.GetUserReply{}
	id, err := s.babyRepo.GeUserId(ctx, req.Id)
	if err != nil {
		return resp, err
	}
	person := &pb.Person{
		Name: "test",
		Age:  18,
	}
	resp = &pb.GetUserReply{
		AppId: cast.ToString(id),
		Name: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"key1": {Kind: &structpb.Value_StringValue{StringValue: "value1"}},
				"key2": {Kind: &structpb.Value_NumberValue{NumberValue: 123.45}},
			},
		},
		Person: person,
		People: &pb.People{
			Persons: []*pb.Person{
				person,
			},
		},
		Human: &pb.Human{
			People: map[string]*pb.Person{
				"key1": person,
			},
		},
		PeopleList: &pb.PeopleListArrMap{
			People: map[string]*pb.People{
				"key1": {
					Persons: []*pb.Person{
						person,
					},
				},
			},
		},
	}
	return resp, nil
}
