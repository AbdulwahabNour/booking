package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/render"
)

 

var Repo *Repository 

type Repository struct{
	App *config.AppConfig
}

func NewRepo(c *config.AppConfig) *Repository{
	return &Repository{
		App: c,
	}
}
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	ipAddress :=r.RemoteAddr
	session, _ := m.App.Session.Get(r, "remote_ip")
	session.Values["ip"] = ipAddress
	fmt.Println(session)
	session.Save(r, w)
	 

    render.RenderTemplate(w, "home.page.tmpl", models.TemplateData{})
	 
	
 }
 
 
 func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	  var dataT string
	  s, _ := m.App.Session.Get(r, "remote_ip") 

	 if  !s.IsNew {
		
		val, ok := s.Values["ip"].(string)
		if ok{
			dataT = val
		}
		fmt.Println(val)

	 }
     dataTemp :=map[string]string{"remote_ip":dataT}

	 fmt.Println("dataOut", dataT== "", dataTemp)

	  data :=  models.TemplateData{
		StringInfo: dataTemp,
	  }
	  fmt.Println(data)

	  render.RenderTemplate(w, "about.page.tmpl",data)
 
 } 

 func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	 
    render.RenderTemplate(w, "make-reservation.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
 

    render.RenderTemplate(w, "generals.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
 

    render.RenderTemplate(w, "majors.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){

    render.RenderTemplate(w, "contact.page.tmpl", models.TemplateData{})
	 
 }
 
 func (m *Repository)SearchAvailability(w http.ResponseWriter, r *http.Request){

    render.RenderTemplate(w, "search-availability.page.tmpl", models.TemplateData{})
	 
 }

 
 func (m *Repository)PostSearchAvailability(w http.ResponseWriter, r *http.Request){

    fmt.Fprint(w,  "POST")
	 
 }

 type jsonResponse struct{
	  OK bool `json:"ok"`
	  Message string `json:"message"`
 }
 func (m *Repository)AvailabilityJson(w http.ResponseWriter, r *http.Request){
	res := jsonResponse{
		OK: true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		log.Println(err)
	}
 
	 w.Write(out)
 }
