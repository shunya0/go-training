package rabbitmq

import (
	"Mongo-GoClient/models"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishVerificationEmail(email string, code string) error {
	var mailVerify models.VerificationEmailEvent
	mailVerify.Code = code
	mailVerify.Email = email
	url := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println("error connecting to queue: ", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err, "ERR")
		return err
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
		return err
	}

	event := mailVerify
	body, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err, "ERR")
		return err
	}
	errss := ch.Publish(
		"",
		queue_name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if errss != nil {
		fmt.Println(err, "ERR")
		return err
	}

	fmt.Println("Email published to queue")
	return nil
}
