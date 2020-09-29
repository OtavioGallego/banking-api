package models

import (
	"regexp"
)

// Account that will be stored in the database
type Account struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

// FormatDocumentNumber removes trailing spaces and non alphanumeric characters
func (account *Account) FormatDocumentNumber() {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	account.DocumentNumber = reg.ReplaceAllString(account.DocumentNumber, "")
}
