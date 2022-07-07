package router

import (
	"github.com/SumitTiket/booking-system/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/hotel", middleware.GetAllHotels).Methods("GET", "OPTIONS")

	return router
}
