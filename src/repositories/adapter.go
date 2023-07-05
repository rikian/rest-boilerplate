package repositories

import (
	"context"
	"golang-test/common"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewRepo(logger *zap.Logger) Repo {
	return &repoImpl{
		logger: logger,
	}
}

type Repo interface {
	Read(context.Context, *gorm.DB, uint) []common.Book
	Create(context.Context, *gorm.DB, *common.Book)
	Update(context.Context, *gorm.DB, *common.Book, string)
	InsertAfterUpdate(context.Context, *gorm.DB, *common.Book, string)
	Delete(context.Context, *gorm.DB, *common.Book, string)

	Rollback(*gorm.DB) error
	Commit(*gorm.DB) error
}

type repoImpl struct {
	logger *zap.Logger
}
