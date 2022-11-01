package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AbdulwahabNour/booking/models"
	"github.com/AbdulwahabNour/booking/pkg/config"
) 

var functions template.FuncMap
var app *config.AppConfig

func NewTemplate(c *config.AppConfig){ 
    app = c
}


func RenderTemplate(w http.ResponseWriter, tmpl string, data models.TemplateData){
 
    t, ok := app.TemplateCache[tmpl]
    if !ok {
        log.Fatalln("template not found")
    }
    buf := new(bytes.Buffer)

    err := t.Execute(buf, data)
    if err != nil{
        fmt.Println("===============================================")
        fmt.Println(err)
        fmt.Println("===============================================")
    }
    buf.WriteTo(w)
   
    
 
} 


 
func CreateTemplateCache()(map[string]*template.Template, error){

    myCache := make(map[string]*template.Template)
  
    
    pages, err := filepath.Glob("../../templates/*.page.tmpl")

    if err != nil{
       return myCache, err 
    } 

    for _, page:= range pages{ 
 
        name := filepath.Base(page)
        
        ts, err  := template.New(name).Funcs(functions).ParseFiles(page,"../../templates/base.layout.tmpl")
       
        
        
        if err !=nil{
            return myCache, err 
        }
        // matches, err := filepath.Glob("../../templates/*.layout.tmpl")
      

        // if err !=nil{  
        //     return myCache, err 
        // }

        // if len(matches) >0 {
            
        //      ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
        //      fmt.Println("ts2 =>", ts)
        //         if err !=nil {  
        //             return myCache, err 
        //            }
        // }   

        myCache[name]=ts
    }


    return myCache,nil
}