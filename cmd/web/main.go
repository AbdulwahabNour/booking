package main

import (
	"log"
	"net/http"

	"github.com/AbdulwahabNour/booking/pkg/config"
	"github.com/AbdulwahabNour/booking/pkg/handlers"
	"github.com/AbdulwahabNour/booking/pkg/render"
	"github.com/gorilla/mux"
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


    r := mux.NewRouter()
  
    r.HandleFunc("/", handlerRepo.Home).Methods("GET")
   
    r.HandleFunc("/about", handlerRepo.About).Methods("GET")
 
    http.ListenAndServe(":8080", r)
}