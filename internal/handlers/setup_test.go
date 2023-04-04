package handlers

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/render"
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
)
var app config.AppConfig
var session *scs.SessionManager
var infolog *log.Logger
var errorlog *log.Logger

func getRoutes() http.Handler{
   gob.Register(models.Reservation{})
   var err error
   session = scs.New()
   session.Lifetime = 3 * time.Hour
   session.IdleTimeout = 20*time.Minute
   session.Cookie.Name= "session"
   session.Cookie.Persist = true
   session.Cookie.SameSite = http.SameSiteLaxMode
   session.Cookie.HttpOnly = true
   session.Cookie.Secure = false
   app.Session = session
   infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) 
   app.InfoLog = infolog
   errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
   app.ErrorLog = errorlog
   render.SetTemplatePath("../../templates/")

    app.TemplateCache, err = render.CreateTemplateCache()
    if err != nil{
         log.Fatal("cannot create template cache")
    }
    render.NewTemplate(&app)
    handlerRepo := NewRepo(&app)
    //set repo in handler page
    NewHandlers(handlerRepo)

 
    r := mux.NewRouter()
  
    r.HandleFunc("/", Repo.Home).Methods("GET")
   
    r.HandleFunc("/about", Repo.About).Methods("GET")
    r.HandleFunc("/generals-quarters", Repo.Generals).Methods("GET")
    r.HandleFunc("/majors-suite", Repo.Majors).Methods("GET")

    r.HandleFunc("/search-availability", Repo.SearchAvailability).Methods("GET")
    r.HandleFunc("/search-availability", Repo.PostSearchAvailability).Methods("POST")
    r.HandleFunc("/search-availability-json", Repo.AvailabilityJson).Methods("POST")


    r.HandleFunc("/make-reservation", Repo.Reservation).Methods("GET")
    r.HandleFunc("/make-reservation", Repo.PostReservation).Methods("POST")
    r.HandleFunc("/reservation-summary", Repo.ReservationSummary)

    
    r.HandleFunc("/contact", Repo.Contact).Methods("GET")
   
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
 
    
    return app.Session.LoadAndSave(r)
    

 
}

 