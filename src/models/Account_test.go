package models_test

import (
	"testing"

	. "github.com/OtavioGallego/banking-api/src/models"
)

type accountNumberTestScenario struct {
	input          string
	expectedOutput string
}

var accountNumberTestScenarios = []accountNumberTestScenario{
	{"709.654.190-02", "70965419002"},
	{"513.215.420-46", "51321542046"},
	{"023.793.710-75", "02379371075"},
	{"635.278.060-16", "63527806016"},
	{" 513.215.420-46", "51321542046"},
	{"54.693.785-X", "54693785X"},
	{"12393213210 ", "12393213210"},
	{"895.491.090-AA ", "895491090AA"},
	{"249.908.580-D8 ", "249908580D8"},
	{"     123.456.789-10     ", "12345678910"},
}

func TestFormatDocumentNumber(t *testing.T) {
	for _, scenario := range accountNumberTestScenarios {
		account := Account{DocumentNumber: scenario.input}
		account.FormatDocumentNumber()

		if account.DocumentNumber != scenario.expectedOutput {
			t.Errorf("Expected %s as formatted number but received %s", scenario.expectedOutput, account.DocumentNumber)
		}
	}
}
