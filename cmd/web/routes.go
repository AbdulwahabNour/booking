package main

import (
	"net/http"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/handlers"

	"github.com/gorilla/mux"
)



func routes(app *config.AppConfig) http.Handler{


    r := mux.NewRouter()
  
    r.HandleFunc("/", handlers.Repo.Home).Methods("GET")
   
    r.HandleFunc("/about", handlers.Repo.About).Methods("GET")
    r.HandleFunc("/generals-quarters", handlers.Repo.Generals).Methods("GET")
    r.HandleFunc("/majors-suite", handlers.Repo.Majors).Methods("GET")

    r.HandleFunc("/search-availability", handlers.Repo.SearchAvailability).Methods("GET")
    r.HandleFunc("/search-availability", handlers.Repo.PostSearchAvailability).Methods("POST")
    r.HandleFunc("/search-availability-json", handlers.Repo.AvailabilityJson).Methods("POST")


    r.HandleFunc("/make-reservation/{from}/{to}/{roomId}", handlers.Repo.Reservation).Methods("GET")
    r.HandleFunc("/make-reservation/{from}/{to}/{roomId}", handlers.Repo.PostReservation).Methods("POST")



    r.HandleFunc("/choose-room/{from}/{to}", handlers.Repo.ChooseRoom).Methods("GET")
    r.HandleFunc("/choose-room/{from}/{to}/{page}", handlers.Repo.ChooseRoom).Methods("GET")


    r.HandleFunc("/rooms", handlers.Repo.Getrooms).Methods("GET")
    r.HandleFunc("/rooms/{page}", handlers.Repo.Getrooms).Methods("GET")
    
    r.HandleFunc("/reservation-summary", handlers.Repo.ReservationSummary)


    
    r.HandleFunc("/contact", handlers.Repo.Contact).Methods("GET")
   
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    return r
} 