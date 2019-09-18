package repository

import (
	"time"
)

type Event struct {
	UUID        int64
	Title       string
	Start       time.Time
	End         time.Time
	Description string
	UserId      uint64    `db:"user_id"`
	NotifyTime  time.Time `db:"notice_time"`
}
