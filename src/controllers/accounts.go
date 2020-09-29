package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/OtavioGallego/banking-api/src/db"
	"github.com/OtavioGallego/banking-api/src/models"
	"github.com/OtavioGallego/banking-api/src/repositories"
	"github.com/OtavioGallego/banking-api/src/responses"
	"github.com/gorilla/mux"
)

// CreateAccount handles the accounts POST endpoint
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var account models.Account
	if err = json.Unmarshal(requestBody, &account); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	account.FormatDocumentNumber()
	if account.DocumentNumber == "" || len(account.DocumentNumber) > 11 {
		responses.Error(w, http.StatusBadRequest, errors.New("The document number must be between 1 and 11 digits"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	account.ID, err = repositories.NewAccountsRepository(db).Create(account)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}

// GetAccount handles the account get by id endpoint
func GetAccount(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	accountID, err := strconv.Atoi(parameters["accountId"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	account, err := repositories.NewAccountsRepository(db).GetByID(accountID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if account.ID == 0 {
		responses.Error(w, http.StatusNotFound, errors.New("Account not found"))
		return
	}
	responses.JSON(w, http.StatusOK, account)
}
