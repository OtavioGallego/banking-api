package routes

import (
	"net/http"

	"github.com/OtavioGallego/banking-api/src/controllers"
)

var transactionRoute = route{
	uri:     "/transactions",
	method:  http.MethodPost,
	handler: controllers.CreateTransaction,
}
