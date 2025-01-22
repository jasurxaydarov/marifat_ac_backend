package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(otp string) error {

	var (
		fromGmail = "xaydarovjasur6999@gmail.com"
		password  = "pzlt skhd nvgq fxks"
	)

	var to []string = []string{"nodirxaydarov911@gmail.com"}
	smsHost := "smtp.gmail.com"
	port := "587"

	message := []byte("user data \n" + otp)

	byteM := []byte(message)

	auth := smtp.PlainAuth("", fromGmail, password, smsHost)

	err := smtp.SendMail(smsHost+":"+port, auth, fromGmail, to, byteM)

	if err != nil {
		log.Println("error", err)
		return err
	}

	fmt.Println("sms successfully")

	return nil
}
