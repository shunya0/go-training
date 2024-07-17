package controllers

import (
	"Mongo-GoClient/models"
	"Mongo-GoClient/services"
	"Mongo-GoClient/utils"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type session struct {
	id     string
	userID string
	expiry int64
}

var Sessions = map[string]session{}

func getUniqueID() string {
	id := make([]byte, 32)
	rand.Read(id)
	return base64.URLEncoding.EncodeToString(id)
}

func createSession(userID string) (string, error) {
	sessionID := getUniqueID()
	expiry := time.Now().Add(utils.SessionDuration * time.Second).Unix()

	Sessions[sessionID] = session{
		id: sessionID, userID: userID, expiry: expiry,
	}
	return sessionID, nil
}

func GetSession(sessionID string) (session, bool) {
	sess, exists := Sessions[sessionID]
	if !exists || sess.expiry < time.Now().Unix() {
		return session{}, false
	}
	return sess, true
}

func invalidSession(sessionID string) {
	delete(Sessions, sessionID)
}

func Login(c *gin.Context) {
	var json models.LoginUser

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})

		return
	}

	userExistsBool, err := services.CheckUserExistsService(json)
	if userExistsBool == false && err != nil {
		fmt.Println("error userExistsBool: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	userValidBool, err := services.UserValid(json)
	if userValidBool == false || err != nil {
		fmt.Println("error userValid: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect email/password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  json.Email,
		"expt": time.Now().Add(utils.SessionDuration * time.Second).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(utils.SECRET_KEY_TOKEN))

	if err != nil {
		fmt.Println("error getting token: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	session_id, err := createSession(tokenString)
	if err != nil {
		fmt.Println("can not create session")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}
	
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_id",
		Value:   session_id,
		Expires: time.Now().Add(utils.SessionDuration * time.Second),
	})
	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(c.Writer, "User logged in successfully")

}

func Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		fmt.Println("can not request cookie")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}
	invalidSession(cookie.Value)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})
	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(c.Writer, "User logged out successfully")
}

func SessionMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		fmt.Println("can not request cookie", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}

	sessionID := cookie.Value
	sess, valid := GetSession(sessionID)
	if !valid {
		fmt.Println("not Authorized")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}
	c.Request.Header.Set("user_id", sess.userID)
	c.Next()

}

func Register(c *gin.Context) {
	var json models.RegisterUser
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("error: json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	json.Role = "customer"
	user_id, err := services.CreateUser(json)
	if err != nil {
		fmt.Println("creating user_id: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with username or mail already exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	fmt.Fprintln(c.Writer, "User created with ID: ", user_id)

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
