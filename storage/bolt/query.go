package bolt

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/alyyousuf7/gocash/transaction"
	"github.com/boltdb/bolt"
)

type Bolt struct {
	dbFile string
	db     *bolt.DB
}

func New(dbFile string) *Bolt {
	return &Bolt{
		dbFile: dbFile,
	}
}

func (b *Bolt) Connect() error {
	db, err := bolt.Open(b.dbFile, 0755, nil)
	if err != nil {
		return err
	}

	b.db = db
	return nil
}

func (b *Bolt) Disconnect() error {
	return b.db.Close()
}

func (b *Bolt) PrepareTransactionBucket() error {
	return b.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Transactions"))
		if err != nil {
			return err
		}

		return nil
	})
}

func (b *Bolt) GetAllTransactions() (transaction.Summary, error) {
	var txs transaction.Summary

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Transactions"))

		return bucket.ForEach(func(k, v []byte) error {
			tx := Transaction{}

			if err := json.Unmarshal(v, &tx); err != nil {
				return err
			}

			txs = append(txs, tx)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return txs, err
}

func (b *Bolt) AddTransaction(t time.Time, amount int, note string) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Transactions"))

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		row := Transaction{
			TxID:     id,
			TxTime:   t.Format("2006-01-02 15:04:05"),
			TxAmount: amount,
			TxNote:   note,
		}

		buf, err := json.Marshal(row)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(strconv.FormatUint(row.TxID, 10)), buf); err != nil {
			return err
		}

		return nil
	})
}

func (b *Bolt) RemoveTransaction(id uint64) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Transactions"))

		if err := bucket.Delete([]byte(strconv.FormatUint(id, 10))); err != nil {
			return err
		}

		return nil
	})
}
