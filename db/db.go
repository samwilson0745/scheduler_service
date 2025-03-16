package db

import (
	"fmt"
	"log"
	"os"
	"scheduler_service/internal/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
	err  error
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		fmt.Println(dsn)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		if err := models.MigrateDB(DB); err != nil {
			log.Fatalf("failed to migrate database: %v", err.Error())
		}
		log.Println("Connected to PostgreSQL")
	})
	return DB
}
