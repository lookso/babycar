package tree

import (
	pb "babycare/api/tree/v1"
	"babycare/internal/biz/tree"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TreeService struct {
	pb.UnimplementedTreeServer
	treeBiz *tree.TreeBiz
	log     *log.Helper
}

func NewTreeService(treeBiz *tree.TreeBiz, logger log.Logger) *TreeService {
	return &TreeService{
		treeBiz: treeBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/api")),
	}
}
func (s *TreeService) GetTree(ctx context.Context, req *pb.GetTreeRequest) (*pb.GetTreeReply, error) {
	return s.treeBiz.GetTree(ctx, req)
}
