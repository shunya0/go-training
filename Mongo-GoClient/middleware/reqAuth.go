package middleware

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {

	sess_token, err := c.Request.Cookie("session_id")
	if err != nil {
		fmt.Println("cookie req , auth: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "something went wrong",
		})
		return
	}
	token_str := string(sess_token.Value)

	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SECRET_KEY_TOKEN), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.LoginUser
		//Finding user with token
		c.Set("user", user)
		c.Next()

	} else {
		fmt.Println("auth err", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
