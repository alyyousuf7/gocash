package sqlite

import (
	"time"
	"log"
	"fmt"
)

type Transaction struct {
	id     int
	time   string
	amount int
	note   string
}

func (t Transaction) ID() string {
	return fmt.Sprintf("%d", t.id)
}

func (t Transaction) Time() time.Time {
	newTime, err := time.Parse("2006-01-02T15:04:05Z", t.time)
	if err != nil {
		log.Printf("Invalid time: %s", t.time)
		return time.Time{}
	}
	return newTime
}

func (t Transaction) Amount() int {
	return t.amount
}

func (t Transaction) Note() string {
	return t.note
}
