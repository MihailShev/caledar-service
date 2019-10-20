package main

import (
	"encoding/json"
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/MihailShev/calendar-service/pkg/connector"
	"github.com/MihailShev/calendar-service/services/scaner/db"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Config struct {
	DB string
	AMQP
}

type AMQP struct {
	Addr            string
	NotifyQueue     string
	NotifyExchange  string
	ConnectionDelay time.Duration
	ConnectionTry   int
}

type Event struct {
	UUID        int64
	Title       string
	Start       time.Time
	End         time.Time
	NotifyTime  time.Time `db:"notice_time"`
	Description string
	UserId      uint64 `db:"user_id"`
}

type Queue struct {
	q        amqp.Queue
	ch       *amqp.Channel
	exchange string
}

func main() {
	var config = Config{}
	err := conf.Read("./", &config)

	failOnError(err)

	scan, err := db.NewEventScanner(connector.Config{Dns: config.DB})

	failOnError(err)

	forever := make(chan struct{})

	go func() {
		q := getQueue(config.AMQP)
		ticker := time.NewTicker(10 * time.Second)
		lastScan := time.Now().Add(-10 * time.Second)

		for {
			curTime := time.Now()
			sendEvents(scan, q, lastScan, curTime)
			lastScan = curTime

			<-ticker.C
		}
	}()

	log.Println("Start event scanner")

	<-forever
}

func sendEvents(eventScanner *db.EventScanner, q *Queue, dateFrom time.Time, dateTo time.Time) {
	events, err := eventScanner.Scan(dateFrom, dateTo)

	logError(err)

	for _, v := range events {
		body, err := json.Marshal(v)

		if err != nil {
			logError(err)
			continue
		}

		publish(q, body)
	}
}

func publish(q *Queue, body []byte) {
	err := q.ch.Publish(
		q.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Println("Publish error", err)
	} else {
		log.Println("Publish to notify queue", string(body))
	}

}

func getQueue(conf AMQP) *Queue {
	conn, err := queueConnect(conf)
	failOnError(err)

	ch, err := conn.Channel()
	failOnError(err)

	q, err := ch.QueueDeclare(
		conf.NotifyQueue, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err)

	err = ch.ExchangeDeclare(conf.NotifyExchange,
		"direct", true, false, false, false, nil)
	failOnError(err)

	err = ch.QueueBind(conf.NotifyQueue, "", conf.NotifyExchange, false, nil)
	failOnError(err)

	return &Queue{
		q:        q,
		ch:       ch,
		exchange: conf.NotifyExchange,
	}
}

func queueConnect(config AMQP) (*amqp.Connection, error) {
	connectionTry := 0
	var conn *amqp.Connection
	var err error

	for connectionTry < config.ConnectionTry {
		connectionTry++

		log.Println("Trying connect to queue broker", connectionTry)

		conn, err = amqp.Dial(config.Addr)

		if err == nil {
			log.Println("Connect to queue broker is established")
			break
		}

		log.Println("Failed to connect to queue broker", err)

		time.Sleep(config.ConnectionDelay * time.Second)
	}

	return conn, err
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
