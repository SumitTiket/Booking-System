package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SumitTiket/booking-system/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server started on PORT: 8080..")

	log.Fatal(http.ListenAndServe(":8080", r))
}
