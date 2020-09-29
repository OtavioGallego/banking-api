package repositories

import (
	"database/sql"

	"github.com/OtavioGallego/banking-api/src/models"
)

// NewAccountsRepository creates a new accounts repository and returns it
func NewAccountsRepository(db *sql.DB) *Accounts {
	return &Accounts{db}
}

// Accounts represents an accounts repository
type Accounts struct {
	db *sql.DB
}

// Create adds an account into the database
func (repo Accounts) Create(account models.Account) (int, error) {
	statement, err := repo.db.Prepare(
		"insert into accounts (document_number) values(?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(account.DocumentNumber)
	if err != nil {
		return 0, err
	}

	lastInsertID, _ := result.LastInsertId()
	return int(lastInsertID), nil
}

// GetByID retrieves one account from the database
func (repo Accounts) GetByID(ID int) (models.Account, error) {
	rows, err := repo.db.Query("select * from accounts where id = ?", ID)
	if err != nil {
		return models.Account{}, err
	}
	defer rows.Close()

	var account models.Account
	if rows.Next() {
		if err = rows.Scan(&account.ID, &account.DocumentNumber); err != nil {
			return models.Account{}, err
		}
	}

	return account, nil
}
