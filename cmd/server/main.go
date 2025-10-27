package main

import (
	"log"
	"context"
	"database/sql"

	"task/config"
	db "task/db/sqlc"
	"task/internal/handler"
	"task/internal/middleware"
	"task/internal/repository"
	"task/internal/routes"
	"task/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	cfg := config.LoadConfig()
	log.Printf(" Connected config loaded. DB URL: %s", cfg.DatabaseURL)

	// // Connect to Postgres
	// dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	// if err != nil {
	// 	log.Fatalf("failed to connect to db: %v", err)
	// }
	// defer dbpool.Close()

	// Connect to Postgres using database/sql with pgx driver
    dbpool, err := sql.Open("pgx", cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("failed to connect to db: %v", err)
    }
    defer dbpool.Close()

	// Ping to verify connection
    if err := dbpool.PingContext(context.Background()); err != nil {
        log.Fatalf("failed to ping db: %v", err)
    }

	// Initialize SQLC queries
	q := db.New(dbpool)

	// Repository + Service + Handler
	repo := repository.NewUserRepository(q)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)


	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "User Management API",
	})


	// Register middleware
	app.Use(middleware.RequestLogger())
	

	// Health check route (helps test server startup)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running smoothly",
		})
	})

	// Register user routes
	routes.RegisterUserRoutes(app, h)

	// Start the server on port 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
