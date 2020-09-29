package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OtavioGallego/banking-api/src/router"
)

func main() {
	fmt.Println("API Listening on port 5000")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))
}
