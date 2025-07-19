package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Hyperion147/gometaverse/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3001")
	user := getEnv("DB_USER", "hyperion")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "gometaverse")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	
	log.Println("Database connection successfull", DB)

}

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.Avatar{},
		&models.User{},
		&models.Element{},
		&models.Space{},
		&models.SpaceElement{},
	)

	if err != nil {
		log.Fatal("Failed to migrate", err)
	}

	log.Println("Database migrated")

}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
