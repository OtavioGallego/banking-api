package repositories_test

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OtavioGallego/banking-api/src/models"
	. "github.com/OtavioGallego/banking-api/src/repositories"
)

func TestCreateTransaction(t *testing.T) {
	transaction := models.Transaction{AccountID: 1, OperationTypeID: 3, Amount: -150, EventDate: time.Now()}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error while opening database connection: %s", err)
	}
	defer db.Close()

	t.Run("Success", func(t *testing.T) {
		mock.ExpectPrepare("insert into transactions (.+) values").ExpectExec().
			WithArgs(transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate).
			WillReturnResult(sqlmock.NewResult(1, 1))

		transactionID, err := NewTransactionsRepository(db).Create(transaction)

		if transactionID != 1 {
			t.Errorf("Expected transaction id 1 but received %d", transactionID)
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
		mock.ExpectPrepare("insert into transactions (.+) values").WillReturnError(prepareError)

		transactionID, err := NewTransactionsRepository(db).Create(transaction)
		if transactionID != 0 {
			t.Errorf("Expected account id 0 but received %d", transactionID)
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

		mock.ExpectPrepare("insert into transactions (.+) values").ExpectExec().
			WithArgs(transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate).
			WillReturnError(execError)

		transactionID, err := NewTransactionsRepository(db).Create(transaction)
		if transactionID != 0 {
			t.Errorf("Expected account id 0 but received %d", transactionID)
		}

		if err == nil {
			t.Errorf("Expected to receive error %s", execError)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("There were unfulfilled expectations: %s", err)
		}
	})
}
