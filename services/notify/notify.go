package main

import (
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/streadway/amqp"
	"log"
)

type Config struct {
	Addr        string
	NotifyQueue string
}

func main() {
	var config = Config{}
	err := conf.Read("./", &config)

	if err != nil {
		failOnError(err, "Failed to read config")
	}

	conn, err := amqp.Dial(config.Addr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.NotifyQueue, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			notify(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}

func notify(m []byte) {
	log.Println("Event notify", string(m))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
