package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var CONNECTION *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Tried to load .env but it failed... Using os env!")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	CONNECTION, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("connection error:", err)
	} else {
		log.Println("Connected to database")
	}
}

func MigrateTables(models ...interface{}) {
	err := CONNECTION.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
}
