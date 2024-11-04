package main

import (
	"github.com/AkselRivera/stori-challenge/migration-service/cmd/api/handlers/health"
	"github.com/AkselRivera/stori-challenge/migration-service/cmd/api/handlers/migration"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/infrastructure/sender"
	"github.com/AkselRivera/stori-challenge/migration-service/pkg/repositories/postgres"
	migrationPostgres "github.com/AkselRivera/stori-challenge/migration-service/pkg/repositories/postgres/migration"
	emailService "github.com/AkselRivera/stori-challenge/migration-service/pkg/services/email"
	migrationService "github.com/AkselRivera/stori-challenge/migration-service/pkg/services/migration"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "github.com/AkselRivera/stori-challenge/migration-service/docs"
)

// @title Migration Service | API Docs
// @version 1.0
// @description This is a migration service for Stori Challenge
// @contact.name API Support
// @contact.email moralesaksel@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Use(logger.New(logger.Config{
		Format: "FIBER - [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Connect to database
	client, err := postgres.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	client.AutoMigrate(&domain.Transaction{})

	// Add Repositories
	migrationRepo := &migrationPostgres.Repository{
		Client: client,
	}

	resendClient := sender.ConnectResend()

	// Add Infraestructure
	emailSender := &sender.ResendEmailSender{
		Client: resendClient,
	}

	// Add Services
	emailSrv := &emailService.Service{
		Sender: *emailSender,
	}

	migrationSrv := &migrationService.Service{
		Repo:   migrationRepo,
		Sender: emailSrv,
	}

	// Add Handlers
	healthHandler := &health.Handler{}
	migrationHandler := &migration.Handler{
		MigrationService: migrationSrv,
	}

	// Swagger docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Add Routes
	app.Get("/health", healthHandler.Check)

	app.Post("/migrate", migrationHandler.Migrate)

	app.Listen(":8080")
}
