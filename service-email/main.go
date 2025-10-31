package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var ServiceSMTP = true

func main() {

	rabbitmq := RabbitMQ{QueueName: "email"}
	rabbitmq.Consume()
}

type RabbitMQ struct {
	Body      string
	QueueName string
}

func (r *RabbitMQ) Consume() {
	for {
		ch, conn := RabbitConnect()
		if ch == nil || conn == nil {
			log.Println(" [!] Gagal koneksi RabbitMQ. Coba lagi dalam 5 detik...")
			time.Sleep(5 * time.Second)
			continue
		}

		err := r.consumeMessages(ch)
		if err != nil {
			return
		}

		RabbitClose(ch, conn)
		time.Sleep(5 * time.Second)
	}
}

type RabbitMessage struct {
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

func (r *RabbitMQ) consumeMessages(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		r.QueueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		var message RabbitMessage
		err := json.Unmarshal(d.Body, &message)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(message.userId, message.email, message.timestamp, message)

		if ServiceSMTP {
			ch.Ack(d.DeliveryTag, true)
		} else {
			ch.Nack(d.DeliveryTag, true, true)
		}
	}

	return nil
}

func RabbitConnect() (*amqp.Channel, *amqp.Connection) {

	connection, err := amqp.Dial("amqp://user:password@127.0.0.1:5672/")
	if err != nil {
		log.Fatal(err)
	}
	// defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	// defer channel.Close()

	return channel, connection
}

func RabbitClose(channel *amqp.Channel, conn *amqp.Connection) {
	defer conn.Close()    //rabbit mq close
	defer channel.Close() //rabbit mq channel close
}
