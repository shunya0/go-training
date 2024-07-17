package main

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type VerificationEmailEvent struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func main() {
	url := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println(err, "ERR")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err, "ERR")
		return
	}

	queue_name := "verification_emails"

	_, errs := ch.QueueDeclare(
		queue_name,
		false,
		false,
		false,
		false,
		nil,
	)
	if errs != nil {
		fmt.Println("error: ", err)
		return
	}

	msgs, err := ch.Consume(
		queue_name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			var event VerificationEmailEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				fmt.Println("can not unmarshel body: ", err)
				continue
			}
			fmt.Println("Received msg:\n", event.Email, "\n", event.Code, "\n")
		}
	}()

	fmt.Println("Waiting for msgs")

	<-forever

}
