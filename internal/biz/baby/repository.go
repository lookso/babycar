package baby

import (
	"babycare/internal/data/model"
	"context"
)

type IBabyRepo interface {
	GeUserId(ctx context.Context, id int64) (int64, error)
	GetStoryList(ctx context.Context, lastId, size int) ([]model.Story, error)
}
