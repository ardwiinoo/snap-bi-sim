package main

import (
	"log"

	"github.com/ardwiinoo/snap-bi-sim/auth-service/cache"
	"github.com/ardwiinoo/snap-bi-sim/auth-service/config"
	"github.com/ardwiinoo/snap-bi-sim/auth-service/database"
	"github.com/ardwiinoo/snap-bi-sim/auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitPostgresDB(cfg)
	rdb := cache.InitRedis(cfg)

	r := gin.Default()
	routes.Init(r, cfg, db, rdb)

	log.Fatal(r.Run(":" + cfg.Port))
}
