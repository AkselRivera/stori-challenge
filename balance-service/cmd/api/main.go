package main

import (
	"log"

	"github.com/AkselRivera/stori-challenge/balance-service/cmd/api/handlers/health"
	"github.com/AkselRivera/stori-challenge/balance-service/cmd/api/handlers/user"
	"github.com/AkselRivera/stori-challenge/balance-service/pkg/repositpories/postgres"
	userPostgres "github.com/AkselRivera/stori-challenge/balance-service/pkg/repositpories/postgres/user"
	userService "github.com/AkselRivera/stori-challenge/balance-service/pkg/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "github.com/AkselRivera/stori-challenge/balance-service/docs"
)

// @title Balance Service | API Docs
// @version 1.0
// @description This is a balance service for Stori Challenge
// @contact.name API Support
// @contact.email moralesaksel@gmail.com
// @host localhost:8081
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

	// Add Repositories
	userRepo := &userPostgres.Repository{
		Client: client,
	}

	// Add Services
	userSrv := &userService.Service{
		Repo: userRepo,
	}

	// Add Handlers
	healthHandler := &health.Handler{}
	userHandler := &user.Handler{
		UserService: userSrv,
	}

	// Swagger docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Add Routes
	app.Get("/health", healthHandler.Check)

	// Group routes
	userRouter := app.Group("/user")
	userRouter.Get("/:user_id/balance", userHandler.GetBalance)

	app.Listen(":8081")

}
