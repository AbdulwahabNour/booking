package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/forms"
	"github.com/AbdulwahabNour/booking/internal/helper"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/render"
	"github.com/justinas/nosurf"
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
	 
 
	// session.Values["ip"] = ipAddress
	// fmt.Println(session)
	// session.Save(r, w)
	 data := models.TemplateData{}
	 render.AddDefaultData(&data, r)

    render.RenderTemplate(w, r, "home.page.tmpl",data)
 
 
	
 }
 
 
 func (m *Repository) About(w http.ResponseWriter, r *http.Request){
 
	//   s, _ := m.App.Session.Get(r, "remote_ip") 

	//  if  !s.IsNew {
		
	// 	val, ok := s.Values["ip"].(string)
	// 	if ok{
	// 		dataT = val
	// 	}
	// 	fmt.Println(val)

	//  }
    //  dataTemp :=map[string]string{"remote_ip":dataT}

	//  fmt.Println("dataOut", dataT== "", dataTemp)

	//   data :=  models.TemplateData{
	// 	StringInfo: dataTemp,
	//   }
	//   fmt.Println(data)

	  render.RenderTemplate(w, r, "about.page.tmpl",models.TemplateData{})
 
 } 
 
 func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
 
    data := models.TemplateData{
		Form: forms.New(nil),
		CSRFToken: nosurf.Token(r),
	}
	render.AddDefaultData(&data, r)

	render.RenderTemplate(w, r,"make-reservation.page.tmpl", data)
	 
	
 }
 func (m *Repository)PostReservation(w http.ResponseWriter, r *http.Request)  {
	
  
	 err := r.ParseForm()
 
	 if err != nil{
		helper.ServerError(w, err)
		return
	 }
	 form := forms.New(r.PostForm)
	 form.TrimValues()
	 reservation := models.Reservation{
		FirstName: form.Values.Get("first_name"),
		LastName: form.Values.Get("last_name"),
		Email: form.Values.Get("email"),
		Phone: form.Values.Get("phone"),
	 } 
	form.Require("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 5)
	form.MinLength("last_name", 5)
	form.MinLength( "email", 5)
	form.MinLength( "phone", 10)
	form.IsEmail("email")

  
	 
	 if !form.Valid(){
		data := make(map[string]interface{})
		data["reservation"] = reservation 

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", models.TemplateData{
			Form: form,
			Data: data,

		}) 
		return
	 }
 
	
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// s, _ := m.App.Session(r, "reservation")
	// gob.Register(reservation)
	// s.Values["reservation"] = reservation
    // s.Save(r, w)
    // a := s.Store()
    //  a.Save(r, w, s)
	 http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
 }

 func (m *Repository)ReservationSummary(w http.ResponseWriter, r *http.Request){
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok{
		m.App.ErrorLog.Println("can't get error from session")
		m.App.Session.Put(r.Context(), "error", "can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
 
//   reservation, ok := s.Values["reservation"] 

//   if !ok{
// 	fmt.Println("not found")
//   }
//   fmt.Println(reservation, reflect.TypeOf(reservation))
//   data := make(map[string]interface{})
//   data["reservation"] = reservation

	data := make(map[string]interface{})
	data["reservation"]= reservation
	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", models.TemplateData{
	Data:  data,
	})

 }



 func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
 

    render.RenderTemplate(w, r, "generals.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
 

    render.RenderTemplate(w, r, "majors.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){

    render.RenderTemplate(w, r, "contact.page.tmpl", models.TemplateData{})
	 
 }
 
 func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request){

    render.RenderTemplate(w, r, "search-availability.page.tmpl", models.TemplateData{})
	 
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
