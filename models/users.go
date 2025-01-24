package models

import "time"

type UserReq struct {
	User_name    string `json:"username"`
	User_surname string `json:"usersurname"`
	User_email   string `json:"email"`
	Password     string `json:"password"`
	User_number  string `json:"user_number"`
	User_role    string `json:"user_role"`
}
type User struct {
	User_id      string    `json:"user_id"`
	User_name    string    `json:"username"`
	User_surname string    `json:"usersurname"`
	User_email   string    `json:"email"`
	User_number  string    `json:"user_number"`
	Password     string    `json:"password"`
	User_role    string    `json:"user_role"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserIsExists struct{
	UserEmail string `json:"email"`
}

type UserOtpData struct{
	Otp string
	Email string
}