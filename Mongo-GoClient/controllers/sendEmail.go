package controllers

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/rabbitmq"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	from := "gothrowaway28@gmail.com"
	password := "xfhzzfrjckykcvdf"

	subject := "using smtp first time"
	var json models.VerificationEmailEvent
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error , json: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := rabbitmq.PublishVerificationEmail(json.Email, json.Code); err != nil {
		fmt.Println("Verification error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	//SMTP setup
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	conn, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		fmt.Println("error, Dial: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return

	}
	defer conn.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if err := conn.StartTLS(tlsConfig); err != nil {
		fmt.Println("error, tlsConfig: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return

	}

	if err := conn.Auth(auth); err != nil {
		fmt.Println("error, conn.Auth: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := conn.Mail(from); err != nil {
		fmt.Println("error, conn.Mail: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	if err := conn.Rcpt(json.Email); err != nil {
		fmt.Println("error, conn.Rcpt: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	// Send the email body
	w, err := conn.Data()
	if err != nil {
		fmt.Println("error, conn.Data: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	defer w.Close()

	// Write the message to the connection
	_, err = w.Write([]byte("Subject: " + subject + "\r\n" + json.Code))
	if err != nil {
		fmt.Println("error, Write: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	fmt.Fprintln(c.Writer, "Email sent successfully")

}
