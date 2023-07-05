package common

type WebRequestRead struct {
	Offset uint `validate:"number" json:"offset"`
}

type WebRequestCreate struct {
	JudulBuku   string `validate:"required" json:"judul_buku"`
	Description string `validate:"required" json:"description"`
	KategoriID  uint   `validate:"required" json:"kategori"`
	KeywordID   uint   `validate:"required" json:"keyword"`
	Harga       int    `validate:"required" json:"harga"`
	Penerbit    string `validate:"required" json:"penerbit"`
}

type WebRequestUpdate struct {
	ID          string `validate:"required,max=64,uuid" json:"book_id" `
	JudulBuku   string `validate:"required" json:"judul_buku"`
	Description string `validate:"required" json:"description"`
	KategoriID  uint   `validate:"required" json:"kategori"`
	KeywordID   uint   `validate:"required" json:"keyword"`
	Harga       int    `validate:"required" json:"harga"`
	Penerbit    string `validate:"required" json:"penerbit"`
}

type WebRequestDelete struct {
	ID string `validate:"required,max=64,uuid" json:"book_id" `
}
