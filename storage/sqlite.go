package storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alyyousuf7/gocash/storage/sqlite"
	"github.com/alyyousuf7/gocash/transaction"
)

type SQLiteAdapter struct {
	db *sqlite.SQLite
}

func NewSQLiteStore(cfg interface{}) (Adapter, error) {
	if cfg == nil {
		return nil, fmt.Errorf("SQLite configuration not privided")
	}

	config, ok := cfg.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid configuration")
	}

	dbFile, ok := config["file"].(string)
	if !ok {
		return nil, fmt.Errorf("SQLite Database filename missing")
	}

	return &SQLiteAdapter{
		db: sqlite.New(dbFile),
	}, nil
}

func (a *SQLiteAdapter) Connect() error {
	if err := a.db.Connect(); err != nil {
		return err
	}

	return a.db.PrepareTransactionTable()
}

func (a *SQLiteAdapter) Disconnect() error {
	return a.db.Disconnect()
}

func (a *SQLiteAdapter) GetSummary() (transaction.Summary, error) {
	return a.db.GetAllTransactions()
}

func (a *SQLiteAdapter) AddTransaction(t time.Time, amount int, note string) error {
	return a.db.AddTransaction(t, amount, note)
}

func (a *SQLiteAdapter) RemoveTransaction(id string) error {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Invalid ID")
	}
	return a.db.RemoveTransaction(newId)
}
