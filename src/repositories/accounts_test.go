package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OtavioGallego/banking-api/src/models"
	. "github.com/OtavioGallego/banking-api/src/repositories"
)

func TestCreateAccount(t *testing.T) {
	account := models.Account{DocumentNumber: "12345678910"}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error while opening database connection: %s", err)
	}
	defer db.Close()

	t.Run("Success", func(t *testing.T) {
		mock.ExpectPrepare("insert into accounts (.+) values").ExpectExec().
			WithArgs(account.DocumentNumber).WillReturnResult(sqlmock.NewResult(1, 1))

		accountID, err := NewAccountsRepository(db).Create(account)
		if accountID != 1 {
			t.Errorf("Expected account id 1 but received %d", accountID)
		}

		if err != nil {
			t.Errorf("Did not expect error but received %s", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Prepare Error", func(t *testing.T) {
		prepareError := errors.New("Prepare Error")
		mock.ExpectPrepare("insert into accounts (.+) values").WillReturnError(prepareError)

		accountID, err := NewAccountsRepository(db).Create(account)
		if accountID != 0 {
			t.Errorf("Expected account id 0 but received %d", accountID)
		}

		if err == nil {
			t.Errorf("Expected to receive error %s", prepareError)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Exec Error", func(t *testing.T) {
		execError := errors.New("Exec Error")

		mock.ExpectPrepare("insert into accounts (.+) values").ExpectExec().
			WithArgs(account.DocumentNumber).WillReturnError(execError)

		accountID, err := NewAccountsRepository(db).Create(account)
		if accountID != 0 {
			t.Errorf("Expected account id 0 but received %d", accountID)
		}

		if err == nil {
			t.Errorf("Expected to receive error %s", execError)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})
}

func TestGetAccountByID(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error while opening database connection: %s", err)
	}
	defer db.Close()

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "document_number"}).AddRow(1, "12345678910")
		mock.ExpectQuery("select (.+) from accounts where id").WithArgs(1).WillReturnRows(rows)

		account, err := NewAccountsRepository(db).GetByID(1)
		if account.ID != 1 {
			t.Errorf("Expected account id 1 but received %d", account.ID)
		}

		if account.DocumentNumber != "12345678910" {
			t.Errorf("Expected account document number 12345678910 but received %s", account.DocumentNumber)
		}

		if err != nil {
			t.Errorf("Did not expect error but received %s", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Query Error", func(t *testing.T) {
		queryError := errors.New("Query Error")
		mock.ExpectQuery("select (.+) from accounts where id").WithArgs(1).WillReturnError(queryError)

		account, err := NewAccountsRepository(db).GetByID(1)
		if account.ID != 0 {
			t.Errorf("Expected account id 0 but received %d", account.ID)
		}

		if account.DocumentNumber != "" {
			t.Errorf("Expected blank account document number but received %s", account.DocumentNumber)
		}

		if err == nil {
			t.Errorf("Expected to receive error %s", queryError)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Scan Error", func(t *testing.T) {
		scanError := errors.New("Scan error")

		rows := sqlmock.NewRows([]string{"id", "document_number"}).AddRow("Invalid ID", "12345678910")
		mock.ExpectQuery("select (.+) from accounts where id").WithArgs(1).WillReturnRows(rows)

		account, err := NewAccountsRepository(db).GetByID(1)
		if account.ID != 0 {
			t.Errorf("Expected account id 0 but received %d", account.ID)
		}

		if account.DocumentNumber != "" {
			t.Errorf("Expected account document number 12345678910 but received %s", account.DocumentNumber)
		}

		if err == nil {
			t.Errorf("Expected to receive error %s", scanError)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})
}
