package utilities

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB = &gorm.DB{}

func init() {
	fmt.Println("Will connect to the db here.\n")
	dsn := "host=localhost user=postgres dbname=sample_backend port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	PostgresDB = db
}

func GetDB() *gorm.DB {
	return PostgresDB.WithContext(context.Background())
}
