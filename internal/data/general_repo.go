package data

import (
	"babycare/internal/biz/common"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type generalRepo struct {
	data *Data
}

func (g generalRepo) Transaction(ctx context.Context, fc func(tx *gorm.DB) error) error {
	return g.data.db.WithContext(ctx).Transaction(fc)
}

func (g generalRepo) Create(ctx context.Context, v any, scope ...func(tx *gorm.DB) *gorm.DB) (int64, error) {
	tx := g.data.db.WithContext(ctx).Scopes(scope...).Create(v)
	return tx.RowsAffected, tx.Error
}

func (g generalRepo) CreateInBatches(ctx context.Context, v any, size int) error {
	return g.data.db.WithContext(ctx).CreateInBatches(v, size).Error
}

func (g generalRepo) Get(ctx context.Context, dest schema.Tabler, id uint) error {
	return g.data.db.WithContext(ctx).First(dest, id).Error
}

func (g generalRepo) GetBy(ctx context.Context, dest schema.Tabler, scope common.Scope) error {
	return g.data.db.WithContext(ctx).Scopes(scope).First(dest).Error
}

func (g generalRepo) Find(ctx context.Context, dest any, scope common.Scope) error {
	return g.data.db.WithContext(ctx).Scopes(scope).Find(dest).Error
}

func (g generalRepo) Count(ctx context.Context, m schema.Tabler, scope common.Scope) (count int64, err error) {
	err = g.data.db.WithContext(ctx).Model(m).Scopes(scope).Count(&count).Error
	return
}

func (g generalRepo) FindBy(ctx context.Context, dest any, scope common.Scope, pager common.Pagination) (count int64, err error) {
	limit := int(pager.GetSize())
	if limit < 1 {
		limit = 10
	}
	offset := (int(pager.GetPage()) - 1) * limit
	if offset < 0 {
		offset = 0
	}

	err = g.data.db.WithContext(ctx).Model(dest).Scopes(scope).Count(&count).Offset(offset).Limit(limit).Find(dest).Error
	return
}

func (g generalRepo) Update(ctx context.Context, m schema.Tabler, id uint, values map[string]any) error {
	return g.data.db.WithContext(ctx).Model(m).Where("id = ?", id).Updates(values).Error
}

func (g generalRepo) UpdateBy(ctx context.Context, m schema.Tabler, scope common.Scope, values map[string]any) error {
	return g.data.db.WithContext(ctx).Model(m).Scopes(scope).Updates(values).Error
}

func (g generalRepo) Delete(ctx context.Context, m schema.Tabler, id uint) error {
	return g.data.db.WithContext(ctx).Delete(m, id).Error
}

func (g generalRepo) DeleteBy(ctx context.Context, m schema.Tabler, scope common.Scope) error {
	return g.data.db.WithContext(ctx).Scopes(scope).Delete(m).Error
}

func NewGeneralRepo(data *Data) common.GeneralRepo {
	return &generalRepo{
		data: data,
	}
}
