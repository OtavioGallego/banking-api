package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON sets status code and returns a json formatted response
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Error sets status code and returns a json formatted error response
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}
