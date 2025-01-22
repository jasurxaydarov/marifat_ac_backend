package models

import "time"

type UserReq struct {
	User_name   string `json:"user_name"`
	User_email  string `json:"user_email"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type User struct {
	User_id     string    `json:user_id`
	User_name   string    `json:"user_name"`
	User_email  string    `json:"user_email"`
	Password    string    `json:"user_password"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
