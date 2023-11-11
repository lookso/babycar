package data

import (
	"context"
)


func (u *Data) GetTree(ctx context.Context, id int32) (int32, error) {
	if id == 0 {
		id = 100
	}
	return id, nil
}
