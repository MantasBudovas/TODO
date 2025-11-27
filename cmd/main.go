package main

import (
	"app/internal/config"
	"app/internal/module/card"
	"app/pkg/cache"
	"app/pkg/database"
	"app/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DBUrl)
	if err != nil {
		logger.Error("Could not connect to database", err)
		panic(err)
	}

	rdb := cache.Connect(cfg.RedisAddr)

	cardRepo := card.NewRepository(db)
	cardService := card.NewService(cardRepo, rdb)
	cardHandler := card.NewHandler(cardService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		api.POST("/cards", cardHandler.Create)
		api.GET("/cards", cardHandler.GetAll)
		api.POST("/cards/:id", cardHandler.Modify)
	}

	logger.Info("Starting server on " + cfg.ServerPort)
	r.Run(cfg.ServerPort)
}
