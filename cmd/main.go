package main

import (
	"fmt"
	"log"
	config "telego/config"
	"telego/database"
	"telego/models"
	router "telego/router"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}
	server := gin.Default()

	router.SetUpRouter(server, db, cfg)

	port := fmt.Sprintf(":%s", cfg.AppPort)

	fmt.Println("Server is running on ", port)
	err = server.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
