package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/OtavioGallego/banking-api/src/db"
	"github.com/OtavioGallego/banking-api/src/models"
	"github.com/OtavioGallego/banking-api/src/repositories"
	"github.com/OtavioGallego/banking-api/src/responses"
)

// CreateTransaction handles the transaction creation endpoint
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var transaction models.Transaction
	if err = json.Unmarshal(requestBody, &transaction); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = transaction.Validate(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	transaction.EventDate = time.Now()

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	account, err := repositories.NewAccountsRepository(db).GetByID(transaction.AccountID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if account.ID == 0 {
		responses.Error(w, http.StatusForbidden, errors.New("No account found with the informed ID"))
		return
	}

	transaction.ID, err = repositories.NewTransactionsRepository(db).Create(transaction)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, transaction)
}
