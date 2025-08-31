package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/snap-bi-sim/signator-generator/api"
	"github.com/ardwiinoo/snap-bi-sim/signator-generator/internal/config"
	"github.com/ardwiinoo/snap-bi-sim/signator-generator/pkg/signer"
)

func main() {
	cfg := config.LoadConfig()

	svc, err := signer.NewService(cfg.PrivateKeyPath)

	if err != nil {
		log.Fatalf("failed to load private key: %v", err)
	}

	app := fiber.New()

	api.SetupRoutes(app, svc)

	log.Printf("signator-generator running on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}