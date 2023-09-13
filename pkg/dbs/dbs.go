package dbs

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const DATABASE_TIMEOUT = 5 * time.Second

type IDatabase interface {
	GetDB() *gorm.DB
	AutoMigrate(models ...any) error
	WithTransaction(function func() error) error
	Create(ctx context.Context, doc any) error
	CreateInBatches(ctx context.Context, docs any, batchSize int) error
	Update(ctx context.Context, doc any) error
	Delete(ctx context.Context, value any, opts ...FindOption) error
}
