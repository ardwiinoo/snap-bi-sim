package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Route) initAuth() {
	auth := r.v1.Group("/auth")
	{
		auth.POST("/token", r.generateToken)
	}
}

func (r *Route) generateToken(c *gin.Context) {
	// sementara dummy dulu
	c.JSON(http.StatusOK, gin.H{
		"access_token": "dummy-token",
		"expires_in":   900,
	})
}
