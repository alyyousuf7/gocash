package sqlite

import (
	"database/sql"
	"time"

	"github.com/alyyousuf7/gocash/transaction"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	dbFile string
	db     *sql.DB
}

func New(dbFile string) *SQLite {
	return &SQLite{
		dbFile: dbFile,
	}
}

func (s *SQLite) Connect() error {
	db, err := sql.Open("sqlite3", s.dbFile)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *SQLite) Disconnect() error {
	return s.db.Close()
}

func (s *SQLite) PrepareTransactionTable() error {
	sqlStmt := `CREATE TABLE IF NOT EXISTS transactions (
		id integer not null primary key,
		dt datetime not null,
		amount integer not null,
		note text
	);`

	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLite) GetAllTransactions() (transaction.Summary, error) {
	sqlQry := `SELECT id, dt, amount, note FROM transactions`

	rows, err := s.db.Query(sqlQry)
	if err != nil {
		return nil, err
	}

	var txs transaction.Summary
	for rows.Next() {
		tx := Transaction{}

		if err := rows.Scan(&tx.id, &tx.time, &tx.amount, &tx.note); err != nil {
			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}

func (s *SQLite) AddTransaction(t time.Time, amount int, note string) error {
	sqlQry := `INSERT INTO transactions (dt, amount, note) VALUES (?, ?, ?)`

	stmt, err := s.db.Prepare(sqlQry)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(t.Format("2006-01-02 15:04:05Z"), amount, note); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) RemoveTransaction(id int) error {
	sqlQry := `DELETE FROM transactions WHERE id = ?`

	stmt, err := s.db.Prepare(sqlQry)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
