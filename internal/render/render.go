package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/justinas/nosurf"
) 

var functions template.FuncMap

var app *config.AppConfig

var templatePath ="templates/"

func SetTemplatePath(name string){
    templatePath = name
}

func NewTemplate(c *config.AppConfig){ 
    app = c
}


func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
    td.Flash = app.Session.PopString(r.Context(),"flash")
    td.Warning= app.Session.PopString(r.Context(), "warning")
    td.Error = app.Session.PopString(r.Context(), "error")
    td.CSRFToken = nosurf.Token(r)
    return td
 
} 


func RenderTemplate(w http.ResponseWriter,r *http.Request,  tmpl string, data models.TemplateData)error{
 
    t, ok := app.TemplateCache[tmpl]
    if !ok {
 
        return errors.New("can't get template from cache")
    }

    buf := new(bytes.Buffer) 
 
    err := t.Execute(buf, data)

    if err != nil{
        return err
    }
    _, err = buf.WriteTo(w)
 
    if err != nil{
        return err
    }
   
    
    return nil
} 


 
func CreateTemplateCache()(map[string]*template.Template, error){

    myCache := make(map[string]*template.Template)
  
    
    pages, err := filepath.Glob(fmt.Sprintf("%s*.page.tmpl",templatePath))

    if err != nil{
       return myCache, err 
    } 

    for _, page:= range pages{ 
 
        name := filepath.Base(page)
        
        ts, err  := template.New(name).Funcs(functions).ParseFiles(page,fmt.Sprintf("%sbase.layout.tmpl",templatePath))
       
        
        
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