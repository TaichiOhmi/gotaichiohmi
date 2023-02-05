package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TaichiOhmi/gotaichiohmi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	DB *gorm.DB
}

var DB Db

func Connect() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("failed to connect to database.", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	err = db.AutoMigrate(&models.Fact{})
	if err != nil {
		log.Fatalln(err)
	}

	DB = Db{
		DB: db,
	}
}
