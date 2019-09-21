package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
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

	db, err := dbConnect()

	failOnError(err)

	forever := make(chan struct{})

	go func() {
		q := getQueue()
		ticker := time.NewTicker(1 * time.Minute)
		lastScan := time.Now()

		sendEvents(db, q, lastScan.Add(-1*time.Minute), lastScan)

		for {
			<-ticker.C

			curTime := time.Now()
			sendEvents(db, q, lastScan, curTime)
			lastScan = curTime
		}
	}()

	<-forever
}

func sendEvents(db *sqlx.DB, q *Queue, dateFrom time.Time, dateTo time.Time) {
	events, err := getEvents(db, dateFrom, dateTo)

	logError(err)

	for _, v := range events {
		body, err := json.Marshal(v)

		if err == nil {
			err = q.ch.Publish(
				"",
				q.q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
		} else {
			logError(err)
		}
	}

	for _, v := range events {
		fmt.Printf("Send to queue %+v\n", v)
	}
}

func dbConnect() (*sqlx.DB, error) {
	dns := "postgres://mshev:123qwe@localhost:5432/calendar?sslmode=disable"
	db, err := sqlx.Open("pgx", dns)

	if err != nil {
		return db, err
	}

	err = db.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}

func getQueue() *Queue {
	conn, err := amqp.Dial("amqp://mshev:123@localhost:5672/")
	failOnError(err)

	ch, err := conn.Channel()
	failOnError(err)

	q, err := ch.QueueDeclare(
		"notify", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err)

	return &Queue{
		q:  q,
		ch: ch,
	}
}

func getEvents(db *sqlx.DB, dateFrom time.Time, dateTo time.Time) ([]Event, error) {
	eventList := make([]Event, 0)
	query := `SELECT * FROM EVENT WHERE notice_time between :dateFrom AND :dateTo`

	rows, err := db.NamedQuery(query, map[string]interface{}{
		"dateFrom": dateFrom,
		"dateTo":   dateTo,
	})

	if err != nil {
		return eventList, err
	}

	for rows.Next() {
		var event Event
		err = rows.StructScan(&event)

		if err != nil {
			return eventList, err
		}

		eventList = append(eventList, event)
	}

	return eventList, nil
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
