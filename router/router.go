package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"telego/config"
	authHandler "telego/handler/auth"
)

func SetUpRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	api := r.Group("/api")

	api.Use(gin.Recovery())

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
	}

}
