package data

import (
	carBiz "babycare/internal/biz/car"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CarData struct {
	data   *Data
	logger *log.Helper
}

func NewCarData(data *Data, logger log.Logger) carBiz.ICarRepo {
	return &CarData{data: data, logger: log.NewHelper(log.With(logger, "module", "data/car"))}
}

// 对user.IUser接口的实现
func (u *CarData) GetUser(ctx context.Context, id int64) (int64, error) {
	return id, nil
}
func (u *CarData) GetAll(ctx context.Context) ([]*carBiz.CarBiz, error) {
	return []*carBiz.CarBiz{}, nil
}

func (u *CarData) Create(ctx context.Context, nickname, phone, pwd string) error {
	return nil
}
