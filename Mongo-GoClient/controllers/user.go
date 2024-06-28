package controllers

import (
	"Mongo-GoClient/utils"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func getSession(sessionID string) (session, bool) {
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
	customer_id_str := c.Request.URL.Query().Get("customer_id")
	utils.CUSTOMER_LOGGED = customer_id_str
	session_id, err := createSession(customer_id_str)
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
	sess, valid := getSession(sessionID)
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
