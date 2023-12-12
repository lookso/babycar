package biz

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Scope func(tx *gorm.DB) *gorm.DB

func IgnoreConflict(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.OnConflict{DoNothing: true}) // 重复的数据忽略
}

type GeneralRepo interface {
	Create(ctx context.Context, v any, scope ...func(tx *gorm.DB) *gorm.DB) error
	CreateInBatches(ctx context.Context, v any, size int) error

	Get(ctx context.Context, dest schema.Tabler, id uint) error
	GetBy(ctx context.Context, dest schema.Tabler, scope Scope) error

	Find(ctx context.Context, dest any, scope Scope) error
	Count(ctx context.Context, m schema.Tabler, scope Scope) (int64, error)
	FindBy(ctx context.Context, dest any, scope Scope, pager Pagination) (int64, error)

	Update(ctx context.Context, m schema.Tabler, id uint, values map[string]any) error
	UpdateBy(ctx context.Context, m schema.Tabler, scope Scope, values map[string]any) error

	Delete(ctx context.Context, m schema.Tabler, id uint) error
	DeleteBy(ctx context.Context, m schema.Tabler, scope Scope) error

	Transaction(ctx context.Context, fc func(tx *gorm.DB) error) error
}

type Pagination interface {
	GetPage() int32
	GetSize() int32
}

func NewPager(page, size int32) Pagination {
	return &pagination{
		page: page,
		size: size,
	}
}

type pagination struct {
	page int32
	size int32
}

func (p *pagination) GetPage() int32 {
	return p.page
}

func (p *pagination) GetSize() int32 {
	return p.size
}
