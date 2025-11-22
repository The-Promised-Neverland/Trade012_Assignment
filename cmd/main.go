package main

import (
	"os"

	"github.com/ThePromisedNeverland/021trade/internal/config"
	"github.com/ThePromisedNeverland/021trade/internal/cron"
	"github.com/ThePromisedNeverland/021trade/internal/db"
	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/repository"
	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/ThePromisedNeverland/021trade/internal/transport/handlers"
	"github.com/ThePromisedNeverland/021trade/internal/transport/router"
)

func main() {
	log := logger.NewLogger()

	app := config.LoadEnv()

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Infof("Connected to DB: %s", os.Getenv("DB_NAME"))
	if err := db.RunMigrations(database); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Infof("Migrations have been completed successfully")

	userRepo := repository.NewUserRepo(database, log)
	rewardRepo := repository.NewRewardRepo(database, log)
	ledgerRepo := repository.NewLedgerRepo(database, log)
	priceRepo := repository.NewPriceRepo(database, log)

	rewardService := services.NewRewardService(rewardRepo, ledgerRepo, priceRepo, log)
	userService := services.NewUserService(userRepo, log)
	ledgerService := services.NewLedgerService(ledgerRepo, log)
	priceService := services.NewPriceService(priceRepo, log)

	userHandler := handlers.NewUserHandler(userService, log)
	rewardHandler := handlers.NewRewardHandler(rewardService, userService, log)
	ledgerHandler := handlers.NewLedgerHandler(ledgerService, log)

	r := router.SetupRouter(userHandler, rewardHandler, ledgerHandler, userService)

	priceCron := cron.NewPriceUpdater(priceRepo, log)
	priceCron.StartHourlyCron()

	portfolioCron := cron.NewPortfolioCron(userService, rewardService, priceService, log)
	portfolioCron.StartMidnightlyCron()

	port := app.Port
	if port == "" {
		port = "8080"
	}
	log.Infof("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
	priceCron.Stop()
}
