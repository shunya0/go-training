package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type VerificationEmailEventTwo struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func main() {
	url := "http://localhost:8080/send-email"
	request_body := VerificationEmailEventTwo{Email: "suryanshrohil05@gmail.com", Code: "ThisIsCode"}
	// fmt.Println("hello")
	forever := make(chan bool)

	go func() {
		for {
			body, err := json.Marshal(request_body)
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}

			req.Header.Set("Content-Type", "aplication/json")
			_, errs := http.DefaultClient.Do(req)
			if errs != nil {
				fmt.Println("error: ", err)
				continue
			}
			fmt.Println("request sent: ", time.Now())
		}
	}()
	fmt.Println("Running bulk testing\n")
	<-forever

}
