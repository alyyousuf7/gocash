package storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alyyousuf7/gocash/transaction"
	"github.com/alyyousuf7/gocash/storage/bolt"
)

type BoltAdapter struct {
	db *bolt.Bolt
}

func NewBoltStore(cfg interface{}) (Adapter, error) {
	if cfg == nil {
		return nil, fmt.Errorf("SQLite configuration not privided")
	}

	config, ok := cfg.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid configuration")
	}

	dbFile, ok := config["file"].(string)
	if !ok {
		return nil, fmt.Errorf("Bolt Database filename missing")
	}

	return &BoltAdapter{
		db: bolt.New(dbFile),
	}, nil
}

func (a *BoltAdapter) Connect() error {
	if err := a.db.Connect(); err != nil {
		return err
	}

	return a.db.PrepareTransactionBucket()
}

func (a *BoltAdapter) Disconnect() error {
	return a.db.Disconnect()
}

func (a *BoltAdapter) GetSummary() (transaction.Summary, error) {
	return a.db.GetAllTransactions()
}

func (a *BoltAdapter) AddTransaction(t time.Time, amount int, note string) error {
	return a.db.AddTransaction(t, amount, note)
}

func (a *BoltAdapter) RemoveTransaction(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Invalid ID")
	}

	return a.db.RemoveTransaction(uint64(newId))
}
