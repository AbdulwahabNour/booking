package main

import (
	"log"
	"net/http"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/handlers"
	"github.com/AbdulwahabNour/booking/internal/render"

	"github.com/gorilla/sessions"
)

var app config.AppConfig
 

func main(){
    var err error
    var store = sessions.NewCookieStore([]byte("secret"))
    app.TemplateCache, err = render.CreateTemplateCache()
    if err != nil{
        log.Fatal(err)
    }
    
    app.Session = store
 
  
    render.NewTemplate(&app)
 
    handlerRepo := handlers.NewRepo(&app)
    
    handlers.NewHandlers(handlerRepo)


    r := routes(&app)
 
    http.ListenAndServe(":8080", r)
}