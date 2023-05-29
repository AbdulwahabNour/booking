package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/handlers"
	"github.com/AbdulwahabNour/booking/internal/helper"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/render"
	"github.com/AbdulwahabNour/booking/internal/repository/dbrepo"
	"github.com/alexedwards/scs/v2"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/microcosm-cc/bluemonday"
)

var app config.AppConfig
var session *scs.SessionManager
var infolog *log.Logger
var errorlog *log.Logger

func main(){
  
    db,err := run()


    if err != nil{
        log.Fatalln(err)
    }
    defer db.Close()
    
    r := routes(&app)
   
    fmt.Println("Server runing on port:8080")
    err = http.ListenAndServe(":8080",nosurf.New(app.Session.LoadAndSave(r)))
    if err != nil{
        log.Fatalln(err)
    }
} 


func declareSession( * scs.SessionManager){
    session = scs.New()
    session.Lifetime = 24 * time.Hour
 
    //If you want the session to disappear the moment that someOne Close the browser window or quits the browser
    //you set that default and the session will not persist the next time they open a window 
    session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
    session.Cookie.Secure = false
}

func run() (*sqlx.DB, error){
   
    gob.Register(models.Reservation{})
    gob.Register(models.Restriction{})
    gob.Register(models.Room{})
    gob.Register(models.User{})

    var err error
    app.InProduction = false
    infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) 
    app.InfoLog = infolog
    errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    app.ErrorLog = errorlog
    app.Bluemonday =  bluemonday.NewPolicy()
 

    declareSession(session)

    app.Session = session
    
    app.TemplateCache, err = render.CreateTemplateCache()
    if err != nil{ 
       return nil,err
    }

    log.Println("Connecting to database......")
 
    //db, err := driver.ConnectSql()
    dbx, err := sqlx.Connect("postgres", "host=localhost sslmode=disable port=5432 dbname=bookings user=postgres password= ")
   
    if err != nil{
        log.Fatalf("can't connect to database err %s", err.Error())
    }

    render.SetConfigToRender(&app)
   
    postgresRepo := dbrepo.NewPostgressRepo(dbx, &app)    
    
    handlerRepo := handlers.NewRepo(postgresRepo ,&app)
    
    handlers.NewHandlers(handlerRepo)
    helper.NewHelper(&app)

    return dbx, nil

}