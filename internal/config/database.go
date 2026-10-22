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
		runMigration(db)
		
	}


	// สร้าง admin user แม้ว่าจะไม่ได้ migrate (กรณีที่ตารางมีอยู่แล้ว)
	if err := SeedDatabase(db, config); err != nil {
		log.Printf("Admin seeding failed: %v", err)
	}

	log.Println("Database connected successfully")

	return db
}

// ฟังก์ชันสำหรับ migrate
func runMigration(db *gorm.DB) {
	log.Println("Starting database migration...")

	// Migrate all models
	err := db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}

// ฟังก์ชันสำหรับ migrate แบบ manual (สำหรับ CLI)
func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)

	log.Println("Running manual migration...")

	// Migrate all models
	err := db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	log.Println("Manual migration completed successfully")
	return nil
}