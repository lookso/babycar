package tree

import (
	pb "babycare/api/tree/v1"
	"babycare/internal/data"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TreeBiz struct {
	data *data.Data
	log  *log.Helper
}

func NewTreeBiz(data *data.Data, logger log.Logger) *TreeBiz {
	return &TreeBiz{data: data, log: log.NewHelper(log.With(logger, "module", "biz/tree"))}
}

func (s *TreeBiz) GetTree(ctx context.Context, req *pb.GetTreeRequest) (*pb.GetTreeReply, error) {

	id, err := s.data.GetTree(ctx, req.Id)
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetTree err: %v", err)
		return nil, err
	}
	return &pb.GetTreeReply{Id: id, Tree: "mytree"}, nil
}
