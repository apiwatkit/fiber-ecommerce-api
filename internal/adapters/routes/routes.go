package routes

import (
	"fiber-ecommerce-api/internal/adapters/http/handlers"
	"fiber-ecommerce-api/internal/adapters/http/middleware"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

// Setup Route
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	adminHandler := handlers.NewAdminHandler()
	
	// Swagger
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// API routes
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.Get("/profile", authHandler.GetProfile)

	// Admin only routes
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireRole("admin"))
	admin.Get("/dashboard", adminHandler.GetDashboard)
	admin.Post("/register", authHandler.AdminRegister)
}