package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/marifat_ac_backend/api/handlers"
	"github.com/jasurxaydarov/marifat_ac_backend/redis"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
)

func Api(storage storage.StorageI,cache redis.RedisRepoI) {

	router := gin.Default()

	h := handlers.NewHandlers(storage,cache)

	router.GET("/ping")

	router.POST("/user", h.UserCreate)
	router.POST("/user-check", h.CheckUser)


	router.Run(":8080")
}
