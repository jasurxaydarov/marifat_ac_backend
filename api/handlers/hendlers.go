package handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/marifat_ac_backend/mail"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
)

type Handlers struct {
	storage storage.StorageI
}

func NewHandlers(storage storage.StorageI) Handlers {

	return Handlers{storage: storage}
}

func (h *Handlers) UserCreate(ctx *gin.Context) {

	var req models.UserReq

	ctx.BindJSON(&req)

	resp, err := h.storage.UserRepo().CreateUser(context.Background(), req)

	if err != nil {

		fmt.Println(err)
		ctx.JSON(500, err)
	}

	data := resp.User_name + "  " + resp.User_email + "  " + resp.Description

	mail.SendMail(data)
	fmt.Println(data)
	ctx.JSON(200, resp)

}
