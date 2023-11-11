package data

import (
	babyBiz "babycare/internal/biz/baby"
	"babycare/internal/data/model"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

type BabyData struct {
	data   *Data
	logger *log.Helper
}

func NewBabyData(data *Data, logger log.Logger) babyBiz.IBabyRepo {
	return &BabyData{data: data, logger: log.NewHelper(log.With(logger, "module", "data/baby"))}
}

// 对user.IUser接口的实现
func (u *BabyData) GeUserId(ctx context.Context, id int64) (int64, error) {
	if id < 0 {
		return 0, nil
	} else {
		return 100, nil
	}
	return id, nil
}

func (u *BabyData) GetStoryList(ctx context.Context, lastId, size int) ([]model.Story, error) {
	var storyList []model.Story
	err := u.data.db.WithContext(ctx).Where("id>?", lastId).Limit(size).Find(&storyList).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, err
	}
	fmt.Println(utils.ToString(size))
	var schema schema.Schema
	fmt.Println(schema)

	return storyList, nil
}
