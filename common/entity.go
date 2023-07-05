package common

import (
	"time"
)

type Book struct {
	ID          string `gorm:"not null;unique;primary_key" json:"book_id"`
	JudulBuku   string `gorm:"not null;" json:"judul_buku"`
	Description string `gorm:"not null;" json:"description"`
	KategoriID  uint   `gorm:"not null;" json:"kategori"`
	Kategori    Kategori
	KeywordID   uint `gorm:"not null;" json:"keyword"`
	Keyword     Keyword
	Harga       string    `gorm:"not null;" json:"harga"`
	Penerbit    string    `gorm:"not null;" json:"penerbit"`
	Created     time.Time `gorm:"type:date;size:128;not null" json:"created"`
	Updated     time.Time `gorm:"type:date;size:128;not null" json:"updated"`
}

type Kategori struct {
	ID   uint   `gorm:"size:128;not null;unique;primary_key" json:"kategori_id"`
	Name string `gorm:"not null;" json:"name"`
}

type Keyword struct {
	ID   uint   `gorm:"size:128;not null;unique;primary_key" json:"keyword_id"`
	Name string `gorm:"null;" json:"name"`
}
