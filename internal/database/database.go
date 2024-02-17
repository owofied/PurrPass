package database

// import (
// 	"fmt"
// 	"os"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var Conn *gorm.DB

// type Config struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DbName   string
// }

// func ConnectPostgres() {
// 	config := Config{
// 		Host:     os.Getenv("POSTGRES_HOST"),
// 		Port:     os.Getenv("POSTGRES_PORT"),
// 		User:     os.Getenv("POSTGRES_USER"),
// 		Password: os.Getenv("POSTGRES_PASSWORD"),
// 		DbName:   os.Getenv("POSTGRES_DB"),
// 	}

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DbName, config.Port)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	Conn = db
// }
