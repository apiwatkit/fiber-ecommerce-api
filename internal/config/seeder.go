package config

import (
	"log"

	"fiber-ecommerce-api/internal/adapters/persistence/models"
	"fiber-ecommerce-api/internal/core/domain/entities"
	"fiber-ecommerce-api/pkg/utils"

	"gorm.io/gorm"
)

// SeedAdminUser สร้าง admin user ตัวแรกถ้ายังไม่มี
func SeedAdminUser(db *gorm.DB, config *Config) error {

	// ตรวจสอบว่ามี admin user อยู่แล้วหรือไม่
	var count int64
	db.Model(&models.User{}).Where("role = ?", entities.RoleAdmin).Count(&count)

	if count > 0 {
		log.Println("Admin user already exists, skipping seeding")
		return nil
	}

	// ตรวจสอบว่ามีการตั้งค่า admin credentials หรือไม่
	if config.AdminEmail == "" {
		log.Println("⚠️  ADMIN_EMAIL not set, skipping admin user seeding")
		log.Println("💡 To create admin user, set ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_FIRST_NAME, ADMIN_LAST_NAME in .env")
		return nil
	}

	if config.AdminPassword == "" {
		log.Println("⚠️  ADMIN_PASSWORD not set, skipping admin user seeding")
		return nil
	}

	if config.AdminFirstName == "" {
		log.Println("⚠️  ADMIN_FIRST_NAME not set, skipping admin user seeding")
		return nil
	}

	if config.AdminLastName == "" {
		log.Println("⚠️  ADMIN_LAST_NAME not set, skipping admin user seeding")
		return nil
	}

	// ตรวจสอบความซับซ้อนของรหัสผ่าน admin
	if err := utils.ValidatePassword(config.AdminPassword); err != nil {
		log.Printf("⚠️  ADMIN_PASSWORD validation failed: %v", err)
		log.Println("💡 Admin password must contain at least 8 characters with uppercase, lowercase, number, and special character")
		return nil
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(config.AdminPassword)
	if err != nil {
		log.Printf("❌ Error hashing admin password: %v", err)
		return err
	}

	// สร้าง admin user
	adminUser := &models.User{
		Email:     config.AdminEmail,
		Password:  hashedPassword,
		FirstName: config.AdminFirstName,
		LastName:  config.AdminLastName,
		Role:      entities.RoleAdmin,
		IsActive:  true,
	}

	if err := db.Create(adminUser).Error; err != nil {
		log.Printf("❌ Error creating admin user: %v", err)
		return err
	}

	log.Printf("✅ Admin user created successfully: %s", config.AdminEmail)
	log.Printf("👤 Name: %s %s", config.AdminFirstName, config.AdminLastName)
	log.Println("⚠️  Please ensure you're using a secure password!")

	return nil
}