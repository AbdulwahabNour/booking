package config

import (
	"html/template"
	"log"

	"github.com/gorilla/sessions"
)

var App AppConfig
 func NewConfig(){

 }
type AppConfig struct{
     UseCache bool 
     TemplateCache map[string]*template.Template
     InfoLog *log.Logger
     InProduction bool
     Session *sessions.CookieStore
     
   
} 