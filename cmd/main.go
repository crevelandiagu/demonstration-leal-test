package main

import (
	"context"
	"demonstration-leal-test/internal/adapter/config"
	"demonstration-leal-test/internal/adapter/handler"
	"demonstration-leal-test/internal/adapter/repository/postgres"
	"demonstration-leal-test/internal/adapter/route"
	"demonstration-leal-test/internal/core/service"
	"fmt"

	"os"
)

// @title Leal-test
// @version         0.1.0
// @contact.name   Cristian Velandia
func main() {

	// Load environment variables
	config, err := config.New()
	if err != nil {
		fmt.Sprintf("Error loading environment variables", "error", err)
		os.Exit(14)
	}

	// Set logger
	//logger.Set(config.App)

	fmt.Sprintf("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		fmt.Sprintf("Error initializing database connection", "error", err)
		os.Exit(15)
	}
	defer db.Close()

	fmt.Sprintf("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	//err = db.Migrate()
	//if err != nil {
	//	slog.Error("Error migrating database", "error", err)
	//	os.Exit(1)
	//}
	//slog.Info("Successfully migrated the database")

	//Dependency
	lealRepo := postgres.NewLealRepository(db)
	lealService := service.NewLealService(lealRepo)
	lealHandler := handler.NewLealHandler(lealService)

	router, err := route.NewRouter(
		*lealHandler,
	)
	if err != nil {
		fmt.Sprintf("Error initializing router", "error", err)
		os.Exit(2)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	fmt.Sprintf("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		fmt.Sprintf("Error starting the HTTP server", "error", err)
		os.Exit(3)
	}

}
