package bolt

import (
	"fmt"
	"log"
	"time"
)

type Transaction struct {
	TxID     uint64 `json:"id"`
	TxTime   string `json:"time"`
	TxAmount int    `json:"amount"`
	TxNote   string `json:"note"`
}

func (t Transaction) ID() string {
	return fmt.Sprintf("%d", t.TxID)
}

func (t Transaction) Time() time.Time {
	newTime, err := time.Parse("2006-01-02 15:04:05", t.TxTime)
	if err != nil {
		log.Printf("Invalid TxTime: %s", t.TxTime)
		return time.Time{}
	}
	return newTime
}

func (t Transaction) Amount() int {
	return t.TxAmount
}

func (t Transaction) Note() string {
	return t.TxNote
}
