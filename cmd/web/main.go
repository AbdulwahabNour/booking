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
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var infolog *log.Logger
var errorlog *log.Logger

func main(){
  
    err := run()
    if err != nil{
        log.Fatalln(err)
    }
    r := routes(&app)
   
    fmt.Println("Server runing on port:8081")
    http.ListenAndServe(":8081",nosurf.New(app.Session.LoadAndSave(r)))
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

func run() error{
   
    gob.Register(models.Reservation{})
    var err error
    app.InProduction = false
    infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) 
    app.InfoLog = infolog
    errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    app.ErrorLog = errorlog

    declareSession(session)

    app.Session = session
    
    app.TemplateCache, err = render.CreateTemplateCache()
    if err != nil{ 
       return err
    }
    
 
  
    render.NewTemplate(&app)
   
 
    handlerRepo := handlers.NewRepo(&app)
    
    handlers.NewHandlers(handlerRepo)
    helper.NewHelper(&app)

    return nil

}