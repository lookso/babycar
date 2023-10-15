package baby

import (
	pb "babycare/api/baby/v1"
	"babycare/internal/biz/baby"
	"github.com/go-kratos/kratos/v2/log"
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
