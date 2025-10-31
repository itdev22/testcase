package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ProduceKeyService(keyName string, payload any) {

	resultRabbit := payload
	body, err := Serialize(resultRabbit)
	if err != nil {
		fmt.Println(err)
	}

	channel, connection := RabbitConnect()
	defer connection.Close()
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		keyName, // name
		true,    // durable
		false,   // auto delete
		false,   // exclusive
		false,   // no wait
		nil,     // args
	)
	if err != nil {
		fmt.Println(err)
	}

	err = channel.Publish(
		"",         // exchange
		queue.Name, // key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
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

func Serialize(msg any) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}
