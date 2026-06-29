package router

import (
	"telego/config"
	adminHandler "telego/handler/admin"
	authHandler "telego/handler/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	authHandler.SetConfig(cfg)

	api := r.Group("/api")

	api.Use(gin.Recovery())

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	admin := api.Group("/admin")
	{
		admin.GET("/users", adminHandler.GetAllUsers)
	}

}
