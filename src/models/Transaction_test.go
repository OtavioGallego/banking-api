package models_test

import (
	"errors"
	"testing"

	. "github.com/OtavioGallego/banking-api/src/models"
)

var (
	zeroAmountError             error = errors.New("Invalid transaction amount. The transaction value must not be zero")
	invalidOperationTypeError   error = errors.New("Invalid operation type. Please inform a number between 1 and 4")
	invalidAccountIdError       error = errors.New("Invalid account ID. The account ID must not be zero")
	expectedNegativeAmountError error = errors.New("Invalid transaction amount. This transaction must have a negative value")
	expectedPositiveAmountError error = errors.New("Invalid transaction amount. This transaction must have a positive value")
)

type transactionTestScenario struct {
	accountID       int
	operationTypeID int
	amount          float64
	expectedError   error
}

var transactionTestScenarios = []transactionTestScenario{
	{1, 1, 0, zeroAmountError},
	{1, 5, 100, invalidOperationTypeError},
	{1, 1, 100, expectedNegativeAmountError},
	{1, 2, 100, expectedNegativeAmountError},
	{1, 3, 100, expectedNegativeAmountError},
	{1, 4, -100, expectedPositiveAmountError},
	{0, 4, -100, invalidAccountIdError},
	{1, 1, -10, nil},
	{1, 2, -20, nil},
	{1, 3, -30, nil},
	{1, 4, 40, nil},
}

func TestValidateTransaction(t *testing.T) {
	for _, scenario := range transactionTestScenarios {

		transaction := Transaction{
			AccountID:       scenario.accountID,
			OperationTypeID: scenario.operationTypeID,
			Amount:          scenario.amount,
		}

		if err := transaction.Validate(); err == nil {

			if scenario.expectedError != nil {
				t.Errorf("Expected error %s but got error %v", scenario.expectedError.Error(), nil)
			}

		} else if err.Error() != scenario.expectedError.Error() {
			t.Errorf("Expected error %s but got error %s", scenario.expectedError, err)
		}
	}
}
