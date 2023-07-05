package repositories

import (
	"context"
	"golang-test/common"

	"gorm.io/gorm"
)

func (a *repoImpl) Read(ctx context.Context, tx *gorm.DB, offset uint) []common.Book {
	var books []common.Book
	r := tx.Preload("Kategori").Preload("Keyword").Offset(int(offset)).Limit(3).Find(&books)
	if r.Error != nil {
		a.logger.Error(r.Error.Error())
		panic(common.Http404{})
	}
	if r.RowsAffected > 3 {
		a.logger.Error("Row affected not equal 3")
		panic(common.Http500{})
	}
	return books
}

func (a *repoImpl) Create(ctx context.Context, tx *gorm.DB, data *common.Book) {
	r := tx.Create(data)
	if r.Error != nil {
		a.logger.Error(r.Error.Error())
		panic(common.Http400{})
	}
	if r.RowsAffected != 1 {
		a.logger.Error("Row affected not equal 1")
		panic(common.Http500{})
	}
}

func (a *repoImpl) Update(ctx context.Context, tx *gorm.DB, book *common.Book, id string) {
	r := tx.Raw("select * from books where id= ? for update", id).
		First(book)
	if r.Error != nil {
		a.logger.Error(r.Error.Error())
		panic(common.Http404{})
	}
	if r.RowsAffected != 1 {
		a.logger.Error("Row affected not equal 1")
		panic(common.Http500{})
	}
}

func (a *repoImpl) InsertAfterUpdate(ctx context.Context, tx *gorm.DB, book *common.Book, id string) {
	r := tx.Updates(book).Where("id=?", id)
	if r.Error != nil {
		a.logger.Error(r.Error.Error())
		panic(common.Http404{})
	}
	if r.RowsAffected != 1 {
		a.logger.Error("Row affected not equal 1")
		panic(common.Http500{})
	}
}

func (a *repoImpl) Delete(ctx context.Context, tx *gorm.DB, book *common.Book, id string) {
	r := tx.Raw("select * from books where id= ? for update", id).
		First(book)
	if r.Error != nil {
		a.logger.Error(r.Error.Error())
		panic(common.Http404{})
	}
	if r.RowsAffected != 1 {
		a.logger.Error("Row affected not equal 1")
		panic(common.Http500{})
	}
	delete := tx.Delete(book).Where("id=?", id)
	if delete.Error != nil {
		a.logger.Error(delete.Error.Error())
		panic(common.Http404{})
	}
	if delete.RowsAffected != 1 {
		a.logger.Error("Row affected not equal 1")
		panic(common.Http500{})
	}
}

func (a *repoImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (a *repoImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}
