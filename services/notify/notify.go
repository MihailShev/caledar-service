package main

import (
	"fmt"
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Addr            string
	NotifyQueue     string
	NotifyExchange  string
	ConnectionDelay time.Duration
	ConnectionTry   int
	Monitoring      string
}

func main() {
	var config = Config{}
	err := conf.Read("./", &config)

	if err != nil {
		failOnError(err, "Failed to read config")
	}

	conn, err := connect(config)
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

	err = ch.QueueBind(config.NotifyQueue, "", config.NotifyExchange, false, nil)
	failOnError(err, "Failed to bind a queue")

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

	notifyMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   "calendar_service",
		Subsystem:   "notify",
		Name:        "rps",
		Help:        "notify per second",
	})

	prometheus.MustRegister(notifyMetric)
	var rps float64 = 0
		notifyMetric.Set(0)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			//notifyMetric.SetToCurrentTime()
			notifyMetric.Set(rps)

			//notifyMetric.Sub(rps)
			//notifyMetric.Set(rps)
			fmt.Println(rps)
			rps = 0
			<-ticker.C
		}
	}()

	go func() {
		for d := range msgs {
			rps++
			notify(d.Body)
		}
	}()

	go func() {
		err := http.ListenAndServe(config.Monitoring, promhttp.Handler())
		log.Println(err)
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}

func connect(config Config) (*amqp.Connection, error) {
	connectionTry := 0
	var conn *amqp.Connection
	var err error

	for connectionTry < config.ConnectionTry {
		connectionTry++

		log.Println("Trying connect to queue broker", connectionTry)

		//conn, err = amqp.Dial(config.Addr)
		conn, err = amqp.Dial("amqp://guest:guest@localhost:5673/")
		if err == nil {
			log.Println("Connect to queue broker is established")
			break
		}

		log.Println("Failed to connect to queue broker", err)

		time.Sleep(config.ConnectionDelay * time.Second)
	}

	return conn, err
}

func notify(m []byte) {
	log.Println("Event notify", string(m))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
