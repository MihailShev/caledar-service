package main

import (
	"encoding/json"
	"github.com/MihailShev/calendar-service/config"
	"github.com/MihailShev/calendar-service/internal/db"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/streadway/amqp"
	"log"
	"time"
)

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
	conf := config.Conf{}
	configuration, err := conf.GetConfig()

	failOnError(err)

	scan, err := db.NewEventScanner(db.Config{Dns: configuration.Dns})

	failOnError(err)

	forever := make(chan struct{})

	go func() {
		q := getQueue(configuration.AMPQ)
		ticker := time.NewTicker(1 * time.Minute)
		lastScan := time.Now()

		sendEvents(scan, q, lastScan.Add(-1*time.Minute), lastScan)

		for {
			<-ticker.C

			curTime := time.Now()
			sendEvents(scan, q, lastScan, curTime)
			lastScan = curTime
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

		if err == nil {
			publish(q, body)
		} else {
			logError(err)
		}
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

func getQueue(conf config.AMPQ) *Queue {
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
