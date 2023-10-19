package baby


import "context"

type IBabyRepo interface {
	GeUserId(ctx context.Context, id int64) (int64,error)
}

