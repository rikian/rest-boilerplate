package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", host, user, pass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	pool, err := db.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	pool.SetMaxIdleConns(5)
	pool.SetMaxOpenConns(20)

	log.Printf("Success create connection to Postgres server at %v://%v:****:%v/%v ...", host, user, dbPort, dbName)
	return db
}
