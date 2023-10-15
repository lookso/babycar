package baby


import "context"

type IBabyRepo interface {
	GetUser(ctx context.Context, id int64) (int64,error)
}

