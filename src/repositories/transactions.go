package repositories

import (
	"database/sql"

	"github.com/OtavioGallego/banking-api/src/models"
)

// Transactions represents a transactions repository
type Transactions struct {
	db *sql.DB
}

// NewTransactionsRepository creates a new transactions repository
func NewTransactionsRepository(db *sql.DB) *Transactions {
	return &Transactions{db}
}

// Create adds an account into the database
func (repo Transactions) Create(transaction models.Transaction) (int, error) {
	statement, err := repo.db.Prepare(
		"insert into transactions (account_id, operation_type_id, amount, event_date) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)
	if err != nil {
		return 0, err
	}

	lastInsertID, _ := result.LastInsertId()
	return int(lastInsertID), nil
}
