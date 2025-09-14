package routes

import (
	"github.com/ardwiinoo/snap-bi-sim/auth-service/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Route struct {
	cfg *config.AppConfig
	db  *gorm.DB
	rdb *redis.Client
	v1  *gin.RouterGroup
}

func Init(r *gin.Engine, cfg *config.AppConfig, db *gorm.DB, rdb *redis.Client) {
	v1 := r.Group("/v1")

	route := &Route{
		cfg: cfg,
		db:  db,
		rdb: rdb,
		v1:  v1,
	}

	route.initAuth()
}
