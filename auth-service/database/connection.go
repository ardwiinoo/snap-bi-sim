package database

import (
	"fmt"
	"log"

	"github.com/ardwiinoo/snap-bi-sim/auth-service/config"
	"github.com/ardwiinoo/snap-bi-sim/auth-service/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(
		&domains.Partner{},
		&domains.Token{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
