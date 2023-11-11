package baby

import (
	pb "babycare/api/baby/v1"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type BabyBiz struct {
	babyRepo IBabyRepo
	log      *log.Helper
}

func NewBabyBiz(babyRepo IBabyRepo, logger log.Logger) *BabyBiz {
	return &BabyBiz{babyRepo: babyRepo, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}

func (s *BabyBiz) GetStoryList(ctx context.Context, lastId, size int) (*pb.GetStoryListReply, error) {
	reply := new(pb.GetStoryListReply)
	storyList, err := s.babyRepo.GetStoryList(ctx, lastId, size)
	if err != nil {
		return reply, err
	}
	stories := make([]*pb.GetStoryListReply_Story, 0, len(storyList))
	for _, story := range storyList {
		stories = append(stories, &pb.GetStoryListReply_Story{
			Id:         int64(story.ID),
			Title:      story.Title,
			Content:    story.Content,
			Status:     int32(story.Status),
			CreateTime: int64(story.CreateTime),
			UpdateTime: int64(story.UpdateTime),
		})
	}
	reply = &pb.GetStoryListReply{
		Stories: stories,
	}
	return reply, nil
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
	// json.Marshal 无法正确处理 Protocol Buffers 的一些特性，如 oneof 字段，或者特殊的 Well-Known Types，如 Timestamp、Duration、Struct 等
	// 将 Person 消息编码为 JSON
	jsonBytes, err := protojson.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return resp, err
	}
	fmt.Println(jsonBytes, string(jsonBytes))

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
