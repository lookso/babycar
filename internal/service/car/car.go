package car

import (
	pb "babycare/api/car/v1"
	"babycare/internal/biz/car"
	"babycare/internal/pkg"
	"bytes"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
)

type CarService struct {
	pb.UnimplementedCarServer
	carBiz *car.CarBiz
	log    *log.Helper
}

func NewCarService(carBiz *car.CarBiz, logger log.Logger) *CarService {
	return &CarService{
		carBiz: carBiz,
		log:    log.NewHelper(log.With(logger, "module", "service/api")),
	}
}
func (s *CarService) HealthCheck(ctx context.Context, structValue *structpb.Value) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Message: "ok",
	}, nil
}
func (s *CarService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfoReply, error) {
	s.carBiz.CreateUser(ctx, req)
	return &pb.UserInfoReply{}, nil
}

func (s *CarService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	id, _ := s.carBiz.GetUser(ctx, req)
	return &pb.GetUserReply{Id: id}, nil
}
func (s *CarService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {

	// 创建一个 SampleMessage 消息，并设置 test_oneof 字段为 name
	msg :=&pb.ListUserRequest{
		UserFilter: &pb.ListUserRequest_Name{
			Name: "Alice",
		},
	}
	b,_:=protojson.Marshal(msg)
	fmt.Println(string(b))

	// 检查 test_oneof 字段的类型
	switch v := msg.UserFilter.(type) {
	case *pb.ListUserRequest_Name:
		fmt.Println("Name:", v.Name)
	case *pb.ListUserRequest_Age:
		fmt.Println("Age:", v.Age)
	case *pb.ListUserRequest_IsActive:
		fmt.Println("IsActive:", v.IsActive)
	default:
		fmt.Println("Unknown type")
	}

	return &pb.ListUserReply{}, nil
}
func (s *CarService) SendJson(ctx context.Context, req *pb.SendJsonRequest) (*pb.SendJsonReply, error) {
	return &pb.SendJsonReply{}, nil
}
func (s *CarService) AuthToken(ctx context.Context, req *pb.AuthTokenRequest) (*pb.AuthTokenReply, error) {
	return &pb.AuthTokenReply{}, nil
}
func (s *CarService) GetWechatContacts(ctx context.Context, req *pb.GetWechatContactsRequest) (*pb.GetWechatContactsReply, error) {
	return &pb.GetWechatContactsReply{}, nil
}

func (s *CarService) DownloadComment(ctx http.Context) error {
	commentList := make([]car.Comment, 0, 10)
	commentList = append(commentList, car.Comment{
		AppName:     "九阳豆浆机",
		TraceId:     "trace_1190gkv2323kls3h213e39831i2",
		SessionId:   "session_10086",
		UserId:      "user_10086",
		TypeChinese: "华为手机meta60",
		CreatedAt:   "2021-01-01 00:00:00",
		Extra:       "this a test",
	})
	f := pkg.NewExcel()
	sheetName := "Sheet1"
	fileName := "comments.xlsx"
	f.File.NewSheet(sheetName)
	// 设置表头
	headers := []string{"应用名称", "trace_id", "session_id", "user_id", "类型", "时间", "详情"}
	f.SetHeader(sheetName, headers)
	for i, comment := range commentList {
		row := i + 2
		f.File.SetCellValue(sheetName, "A"+strconv.Itoa(row), comment.AppName)
		f.File.SetCellValue(sheetName, "B"+strconv.Itoa(row), comment.TraceId)
		f.File.SetCellValue(sheetName, "C"+strconv.Itoa(row), comment.SessionId)
		f.File.SetCellValue(sheetName, "D"+strconv.Itoa(row), comment.UserId)
		f.File.SetCellValue(sheetName, "E"+strconv.Itoa(row), comment.TypeChinese)
		f.File.SetCellValue(sheetName, "F"+strconv.Itoa(row), comment.CreatedAt)
		f.File.SetCellValue(sheetName, "G"+strconv.Itoa(row), comment.Extra)
	}
	buf := new(bytes.Buffer)
	err := f.File.Write(buf)
	if err != nil {
		return err
	}
	if err = f.File.SaveAs(fileName); err != nil {
		return err
	}
	// 写入http响应体
	if err = f.Write(ctx, buf.Bytes(), fileName); err != nil {
		return err
	}
	return nil
}
