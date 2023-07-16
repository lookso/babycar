package car

import "context"

type ICarRepo interface {
	GetUser(ctx context.Context, id int64) (int64,error)
	GetAll(ctx context.Context) ([]*CarBiz, error)
	Create(ctx context.Context, nickName, pwd, phone string) error
}
