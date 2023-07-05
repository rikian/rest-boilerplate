package migrations

import (
	"golang-test/common"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type table struct {
	tableName interface{}
}

func dropTable(db *gorm.DB) {
	var err error
	err = db.Migrator().DropTable(&common.Book{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Migrator().DropTable(&common.Kategori{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Migrator().DropTable(&common.Keyword{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func RunMigration(db *gorm.DB) {
	dropTable(db)

	// registrasi table here.
	// note: table for foregin key, add first
	t := []*table{
		{tableName: &common.Kategori{}},
		{tableName: &common.Keyword{}},
		{tableName: &common.Book{}},
	}

	for _, table := range t {
		err := db.AutoMigrate(table.tableName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	kName := []string{"remaja", "dewasa", "politik", "teknik", "bisnis"}

	for i := 0; i < len(kName); i++ {
		kategory := &common.Kategori{
			ID:   uint(i + 1),
			Name: kName[i],
		}

		result := db.Create(kategory)
		if result.Error != nil {
			log.Fatal(result.Error.Error())
		}
	}

	kKeyword := []string{"terbaru", "viral", "enak", "makanan", "olahraga"}
	for i := 0; i < len(kKeyword); i++ {
		keword := &common.Keyword{
			ID:   uint(i + 1),
			Name: kName[i],
		}

		result := db.Create(keword)
		if result.Error != nil {
			log.Fatal(result.Error.Error())
		}
	}

	books := map[int]*common.Book{
		1: {
			ID:          uuid.New().String(),
			JudulBuku:   "lorem 1",
			Description: "description 1",
			KategoriID:  uint(1),
			KeywordID:   uint(1),
			Harga:       "Rp 50000 ,-",
			Penerbit:    "hello world",
			Created:     time.Now(),
			Updated:     time.Now(),
		},
		2: {
			ID:          uuid.New().String(),
			JudulBuku:   "lorem 2",
			Description: "description 2",
			KategoriID:  uint(2),
			KeywordID:   uint(2),
			Harga:       "Rp 50000 ,-",
			Penerbit:    "hello world",
			Created:     time.Now(),
			Updated:     time.Now(),
		},
		3: {
			ID:          uuid.New().String(),
			JudulBuku:   "lorem 3",
			Description: "description 3",
			KategoriID:  uint(3),
			KeywordID:   uint(3),
			Harga:       "Rp 50000 ,-",
			Penerbit:    "hello world",
			Created:     time.Now(),
			Updated:     time.Now(),
		},
		4: {
			ID:          uuid.New().String(),
			JudulBuku:   "lorem 4",
			Description: "description 4",
			KategoriID:  uint(4),
			KeywordID:   uint(4),
			Harga:       "Rp 50000 ,-",
			Penerbit:    "hello world",
			Created:     time.Now(),
			Updated:     time.Now(),
		},
		5: {
			ID:          uuid.New().String(),
			JudulBuku:   "lorem 5",
			Description: "description 5",
			KategoriID:  uint(5),
			KeywordID:   uint(5),
			Harga:       "Rp 50000 ,-",
			Penerbit:    "hello world",
			Created:     time.Now(),
			Updated:     time.Now(),
		},
	}

	for _, book := range books {
		result := db.Create(book)
		if result.Error != nil {
			log.Fatal(result.Error.Error())
		}
	}

	log.Print("migration success")
}
