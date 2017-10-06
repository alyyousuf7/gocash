package gocash

import (
	"fmt"
	"time"

	"github.com/alyyousuf7/gocash/storage"
	"github.com/alyyousuf7/gocash/transaction"
)

type App struct {
	store storage.Adapter
}

func NewApp(config *Configuration) (*App, error) {
	if config == nil {
		return nil, fmt.Errorf("Configuration missing")
	}

	if config.Storage == "" {
		return nil, fmt.Errorf("Storage undefined")
	}

	store, err := storage.NewStore(config.Storage, config.StorageConfig)
	if err != nil {
		return nil, err
	}

	if err := store.Connect(); err != nil {
		return nil, err
	}

	return &App{store}, nil
}

func (a *App) Close() error {
	return a.store.Disconnect()
}

func (a *App) GetSummary() (transaction.Summary, error) {
	return a.store.GetSummary()
}

func (a *App) AddTransaction(t time.Time, amount int, note string) error {
	return a.store.AddTransaction(t, amount, note)
}

func (a *App) RemoveTransaction(id string) error {
	return a.store.RemoveTransaction(id)
}
