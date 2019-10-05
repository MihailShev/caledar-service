package main

import (
	"encoding/json"
	"github.com/MihailShev/calendar-service/pkg/config"
	"github.com/MihailShev/calendar-service/pkg/connector"
	"github.com/MihailShev/calendar-service/sevices/scaner/db"
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
	Addr        string
	NotifyQueue string
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
	q  amqp.Queue
	ch *amqp.Channel
}

func main() {
	var config = Config{}
	err := conf.Read("../", &config)

	failOnError(err)

	scan, err := db.NewEventScanner(connector.Config{Dns: config.DB})

	failOnError(err)

	forever := make(chan struct{})

	go func() {
		q := getQueue(config.AMQP)
		ticker := time.NewTicker(1 * time.Minute)
		lastScan := time.Now().Add(-1 * time.Minute)

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
		"",
		q.q.Name,
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
	conn, err := amqp.Dial(conf.Addr)
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

	return &Queue{
		q:  q,
		ch: ch,
	}
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
