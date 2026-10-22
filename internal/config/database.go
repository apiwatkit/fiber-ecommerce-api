package config

import (
	"fmt"
	"log"

	"fiber-ecommerce-api/internal/adapters/persistence/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	if config.AutoMigrate {
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		} else {
			log.Println("Successfully to migrate database")
		}
		
	}


	// สร้าง admin user แม้ว่าจะไม่ได้ migrate (กรณีที่ตารางมีอยู่แล้ว)
	if err := SeedAdminUser(db, config); err != nil {
		log.Printf("Admin seeding failed: %v", err)
	}

	log.Println("Database connected successfully")

	return db
}