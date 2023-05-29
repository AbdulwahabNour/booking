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
	"github.com/AbdulwahabNour/booking/internal/repository/dbrepo"
	"github.com/alexedwards/scs"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
var app config.AppConfig
var session *scs.SessionManager
var infolog *log.Logger
var errorlog *log.Logger

func getRoutes() (*sqlx.DB, http.Handler){

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
    render.SetConfigToRender(&app)
    dbx, err := sqlx.Connect("postgres", "host=localhost sslmode=disable port=5432 dbname=bookings user=postgres password= ")
    if err != nil{
        log.Fatalf("can't connect to database err %s", err.Error())
    }
    postgresRepo := dbrepo.NewPostgressRepo(dbx, &app)    


    handlerRepo := NewRepo(postgresRepo, &app)
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
 
    
    return dbx, app.Session.LoadAndSave(r)
    

 
}

 