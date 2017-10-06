package storage

import (
	"fmt"
	"time"

	"github.com/alyyousuf7/gocash/transaction"
)

type Adapter interface {
	Connect() error
	Disconnect() error

	GetSummary() (transaction.Summary, error)
	AddTransaction(time.Time, int, string) error
	RemoveTransaction(string) error
}

func NewStore(name string, cfg interface{}) (Adapter, error) {
	var (
		store Adapter
		err   error
	)

	switch name {
	case "bolt":
		store, err = NewBoltStore(cfg)
		if err != nil {
			return nil, err
		}
	case "sqlite":
		store, err = NewSQLiteStore(cfg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported storage %s", name)
	}

	return store, nil
}
