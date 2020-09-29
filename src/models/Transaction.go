package models

import (
	"errors"
	"time"

	"github.com/OtavioGallego/banking-api/src/utils"
)

// Transaction that will be stored in the database
type Transaction struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"eventDate"`
}

var (
	negativeAmountOperations = []int{1, 2, 3}
	positiveAmountOperations = []int{4}
)

// Validate checks the account ID, operation type and amount
func (transaction Transaction) Validate() error {
	if transaction.AccountID == 0 {
		return errors.New("Invalid account ID. The account ID must not be zero")
	}

	if transaction.Amount == 0 {
		return errors.New("Invalid transaction amount. The transaction value must not be zero")
	}

	if err := transaction.validateOperationType(); err != nil {
		return err
	}

	return nil
}

func (transaction Transaction) validateOperationType() error {
	allOperations := append(negativeAmountOperations, positiveAmountOperations...)

	validOperationType := utils.Find(allOperations, transaction.OperationTypeID)
	if !validOperationType {
		return errors.New("Invalid operation type. Please inform a number between 1 and 4")
	}

	isPositiveAmountOperation := utils.Find(positiveAmountOperations, transaction.OperationTypeID)
	if isPositiveAmountOperation && transaction.Amount < 0 {
		return errors.New("Invalid transaction amount. This transaction must have a positive value")
	}

	isNegativeAmountOperation := utils.Find(negativeAmountOperations, transaction.OperationTypeID)
	if isNegativeAmountOperation && transaction.Amount > 0 {
		return errors.New("Invalid transaction amount. This transaction must have a negative value")
	}

	return nil
}
