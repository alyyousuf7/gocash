package transaction

import (
	"time"
)

const (
	CURRENCY = "Rs."
)

type Transaction interface {
	ID() string
	Time() time.Time
	Amount() int
	Note() string
}
