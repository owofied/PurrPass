package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
