package main

import (
	_ "fiber-ecommerce-api/docs"
	"fiber-ecommerce-api/internal/adapters/http/handlers"
	"fiber-ecommerce-api/internal/adapters/persistence/repositories"
	"fiber-ecommerce-api/internal/adapters/routes"
	"fiber-ecommerce-api/internal/config"
	"fiber-ecommerce-api/internal/core/services"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// load config
	cfg := config.LoadConfig()

	// setup db
	db := config.SetupDatabase(cfg)

	// initial repo
	userRepo := repositories.NewUserRepository(db)

	// initial service
	authService := services.NewAuthService(userRepo)

	// initial handler
	authHandler := handlers.NewAuthHandler(authService)

	// initial fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
