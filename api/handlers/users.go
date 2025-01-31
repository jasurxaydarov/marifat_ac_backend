package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/marifat_ac_backend/mail"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
)

func (h *Handlers) UserCreate(ctx *gin.Context) {

	var req models.UserReq

	ctx.BindJSON(&req)

	resp, err := h.storage.UserRepo().CreateUser(context.Background(), req)

	if err != nil {

		fmt.Println(err)
		ctx.JSON(500, err)
	}

	/*data := fmt.Sprintf(
	`Username:%s   :
	User Phone:%s  :
	User Email:%s  : `,
	resp.User_name,
	resp.User_number,
	resp.User_email)*/

	///mail.SendAdminMail(data)

	//fmt.Println(data)

	ctx.JSON(200, resp)

}

func (h *Handlers) CheckUser(ctx *gin.Context) {

	var reqBody models.UserIsExists

	err := ctx.BindJSON(&reqBody)

	if err != nil {

		fmt.Println("err on BindJSON", err)
		ctx.JSON(400, err)
		return
	}

	isExists, err := h.storage.UserRepo().IsExists(context.Background(), models.IsExists{
		TableName:  "users",
		ClomunName: "email",
		ExpValue:   reqBody.UserEmail,
	})

	if err != nil {

		fmt.Println("err on storage from users is exists ", err)
		ctx.JSON(500, err)

		return
	}

	if isExists.IsExists {
		ctx.JSON(201, models.IsExistsResp{
			IsExists: isExists.IsExists,
			Status:   "sign-in",
		})

		return
	}

	otp := models.UserOtpData{
		Otp:   mail.GenerateOtp(6),
		Email: reqBody.UserEmail,
	}

	otpDataB, err := json.Marshal(&otp)

	if err != nil {

		fmt.Println("err on json.Marshal", err)
		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.UserEmail, string(otpDataB), 120)

	if err != nil {

		fmt.Println("err on redis cache set ", err)
		ctx.JSON(500, err)
		return
	}

	err = mail.SendMail([]string{reqBody.UserEmail}, otp.Otp)

	if err != nil {

		fmt.Println("err on sent otp to user email", err)
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "registr",
		Message: "we sent otp your email",
	})

}

func (h *Handlers) SignUp(ctx *gin.Context) {

	var otpData models.UserOtpData

	var reqBody models.UserReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {

		fmt.Println("err on BindJSON", err)
		ctx.JSON(401, err.Error())
		return
	}

	otpSData,err:=h.cache.GetDell(ctx,reqBody.User_email)

	if err!=nil{
		fmt.Println("h.cache.GetDell",err)
		ctx.JSON(500,err.Error())
		return
	}

	if otpSData==""{

		ctx.JSON(201,"otp code is expired")
		return
	}

	err=json.Unmarshal([]byte(otpSData),&otpData.Otp)

	if otpData.Otp!=reqBody.Otp{

		
		ctx.JSON(405,"incorrect otp code")
		return
	}

}
