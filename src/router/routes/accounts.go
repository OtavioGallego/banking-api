package routes

import (
	"net/http"

	"github.com/OtavioGallego/banking-api/src/controllers"
)

var accountRoutes = []route{
	{
		uri:     "/accounts",
		method:  http.MethodPost,
		handler: controllers.CreateAccount,
	},
	{
		uri:     "/accounts/{accountId}",
		method:  http.MethodGet,
		handler: controllers.GetAccount,
	},
}
