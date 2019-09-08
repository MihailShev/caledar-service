package repository

import (
	"time"
)

type EventModel struct {
	UUID        int64
	Title       string
	Start       time.Time
	End         time.Time
	Description string
	UserId      uint64
	NoticeTime  uint32
}