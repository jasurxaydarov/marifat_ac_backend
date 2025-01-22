package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/marifat_ac_backend/api/handlers"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
)

func Api(storage storage.StorageI) {

	router := gin.Default()

	h := handlers.NewHandlers(storage)

	router.GET("/ping")

	router.POST("/user", h.UserCreate)

	router.Run(":8080")
}
