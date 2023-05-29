package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/AbdulwahabNour/booking/internal/forms"
	"github.com/AbdulwahabNour/booking/internal/helper"
	"github.com/AbdulwahabNour/booking/internal/models"
	"github.com/AbdulwahabNour/booking/internal/pagination"
	"github.com/AbdulwahabNour/booking/internal/render"
	"github.com/AbdulwahabNour/booking/internal/repository"
	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
)

const(

	DATELAYOUT = "2006-01-02"
) 

var Repo *Repository 

type Repository struct{
	App *config.AppConfig
	DBRepo repository.DatabaseRepo
}

func NewRepo(repo repository.DatabaseRepo,c *config.AppConfig) *Repository{
	return &Repository{
		DBRepo: repo,
		App: c,
	 
	}
}
func NewHandlers(r *Repository){ 
	Repo = r
    
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	 

 
	 data := models.TemplateData{}
	 render.AddDefaultData(&data, r)

    render.Template(w, r, "home.page.tmpl",data)
 
 
	
 }
 
 
 func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	data := models.TemplateData{}
	render.AddDefaultData(&data, r)
	data.Form =forms.New(nil) 
	
 

	  render.Template(w, r, "about.page.tmpl", data)
 
 } 

 func (m *Repository)ChooseRoom(w http.ResponseWriter, r *http.Request){

	data := models.TemplateData{}
	render.AddDefaultData(&data, r)
 
	from := mux.Vars(r)["from"]

	to := mux.Vars(r)["to"]
    
	stDate, sterr := time.Parse(DATELAYOUT, from)
	edDate, ederr := time.Parse(DATELAYOUT, to)

	if stDate.Before(time.Now().UTC().Truncate(24 * time.Hour)){
		 
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/search-availability" , http.StatusTemporaryRedirect)
		return
	} 
	
	//set default  date if error happend
	if sterr != nil || ederr != nil{
		stDate = time.Now()
		edDate= stDate.AddDate(0, 0, 7)
	}

	pageReq, exists := mux.Vars(r)["page"]
	page := 1.0
	if exists{
		pasedPage, err := strconv.Atoi(pageReq)
		if err == nil && pasedPage > 0{
			page = float64(pasedPage)
		}	 
	}
    
	countRomms, err := m.DBRepo.CountRooms(r.Context())
	if err != nil {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/ ", http.StatusTemporaryRedirect)
		return
	}

    pagination := pagination.Pagination{
		 TotalItems: float64(countRomms),
		 PageSize: 10.0,
		 CurrentPage: page,
		 BeforeAndAfterPage: 5,
	}
	totalpageNumber := pagination.TotalPages()
   if totalpageNumber < int(page){
	   pagination.CurrentPage = 1.0
   }

   links := pagination.GeneratePageNumbers()

 
   rooms, err :=  m.DBRepo.SearchAvailabilityForRooms(r.Context(), int(pagination.PageSize), int(pagination.Offset()), stDate, edDate)
 

	if err != nil || len(rooms)==0 {
		m.App.Session.Put(r.Context(), "error", "Now availabilty")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
 
   data.Data= make(map[string]interface{})
   data.Data["rooms"] = rooms
   data.Data["links"] = links
   data.Data["From"] = stDate.Format(DATELAYOUT)
   data.Data["To"] = edDate.Format(DATELAYOUT)
   data.Data["pagesNumber"] =  totalpageNumber
   

   render.Template(w, r, "choose-room.page.tmpl", data)
  
	
}

 
 func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
 
    data := models.TemplateData{
		Form: forms.New(nil),
		CSRFToken: nosurf.Token(r),
	}
	render.AddDefaultData(&data, r)


	stDate, err := time.Parse(DATELAYOUT, mux.Vars(r)["from"])
	if err != nil {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/search-availability" , http.StatusTemporaryRedirect)
		return
	}
	

	edDate, err := time.Parse(DATELAYOUT, mux.Vars(r)["to"]) 
	if err != nil {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/search-availability" , http.StatusTemporaryRedirect)
		return
	}
    
	if stDate.Before(time.Now()){
		 
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/search-availability" , http.StatusTemporaryRedirect)
		return
	} 
	

	RoomId, err := strconv.Atoi(mux.Vars(r)["roomId"])
	if err != nil{
		m.App.Session.Put(r.Context(), "error", "please choose the correct room")
		http.Redirect(w, r, "/" , http.StatusTemporaryRedirect) 
		return
	}	 

	 
	room, err := m.DBRepo.GetRoomById(r.Context(), RoomId)
	 
	if err != nil{
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "please choose the correct room")
		http.Redirect(w, r, "/" , http.StatusTemporaryRedirect)
		return
	}

	available, err := m.DBRepo.CheckAvailabilityByDateAndRoom(r.Context(), room.ID, stDate, edDate)
 
	if available || err != nil{
		    m.App.ErrorLog.Println(err, available)
			m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
	}

	data.Data = make(map[string]interface{})
	data.Data["room"] = room
	data.Data["from"] = stDate.Format(DATELAYOUT)
	data.Data["to"] =  edDate.Format(DATELAYOUT)

  

	render.Template(w, r,"make-reservation.page.tmpl", data)
	 
	
 }
 func (m *Repository)PostReservation(w http.ResponseWriter, r *http.Request)  {
	
  
	 err := r.ParseForm()
 
	 if err != nil{
		helper.ServerError(w, err)
		return
	 }

	 form := forms.New(r.PostForm)
	 form.TrimValues()
 
 
  
	 stDate, err := time.Parse(DATELAYOUT, r.PostForm.Get("start_date"))

	 if err != nil {
	    form.Errors.Add("start_date", forms.StartDateFormErr)
	 }
	 
	 edDate, err := time.Parse(DATELAYOUT, r.PostForm.Get("end_date"))
	 if err != nil {
		form.Errors.Add("end_date", forms.EndDateFormErr)
	 }


	
     roomId, err := strconv.Atoi(r.PostForm.Get("room"))

	 if err != nil {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happened try again later")
		http.Redirect(w, r, "/search-availability" , http.StatusTemporaryRedirect)
		return
	 }
 
	room, err := m.DBRepo.GetRoomById(r.Context(), roomId)

	if err != nil{
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "No room available")
		http.Redirect(w, r, "/search-availability", http.StatusTemporaryRedirect)
		return
	}

	 reservation := models.Reservation{

		FirstName: m.App.Bluemonday.Sanitize(form.Values.Get("first_name")),
		LastName: m.App.Bluemonday.Sanitize(form.Values.Get("last_name")),
		Email: m.App.Bluemonday.Sanitize(form.Values.Get("email")),
		Phone: m.App.Bluemonday.Sanitize(form.Values.Get("phone")),
		StartDate: stDate,
		EndDate:   edDate,
		RoomId: roomId,
		
	 }
	 
	
	form.Require("first_name", "last_name", "email", "phone", "start_date", "end_date", "room")
	form.MinLength("first_name", 5)
	form.MinLength("last_name", 5)
	form.MinLength( "email", 5)
	form.MinLength( "phone", 10)
	form.IsEmail("email")
 
	 if !form.Valid(){
       
		data := make(map[string]interface{})
		data["reservation"] = reservation 
		data["room"] = room
		data["from"] = stDate.Format(DATELAYOUT)
		data["to"] =  edDate.Format(DATELAYOUT)
	

		render.Template(w, r, "make-reservation.page.tmpl", models.TemplateData{
			Form: form,
			Data: data,
			CSRFToken: nosurf.Token(r),


		}) 
		return
	 }
	 
	 //insert Reservation
 
	reservationID, err := m.DBRepo.InsertReservation(r.Context(), &reservation)

	if err != nil{
		helper.ServerError(w, err)
		return
	}

	restriction := models.RoomRestriction{
			RoomID: reservation.RoomId,
			StartDate: reservation.StartDate,
			EndDate: reservation.EndDate,
			ReservationID:  reservationID,
			RestrictionID: 1,
	}
	

	err = m.DBRepo.InsertRoomRestrictions(r.Context(), &restriction)
	if err != nil{
		 helper.ServerError(w, err)
		 return
   }                                                    


     m.App.Session.Put(r.Context(), "reservation", reservation)
	 m.App.Session.Put(r.Context(), "room", room)
   
	 http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
 }

 func (m *Repository)ReservationSummary(w http.ResponseWriter, r *http.Request){
	
	reservation, ok := m.App.Session.Pop(r.Context(), "reservation").(models.Reservation)
	
	if !ok{
		m.App.ErrorLog.Println("can't get error from session")
		m.App.Session.Put(r.Context(), "error", "can't get reservation ")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
  

	room, ok := m.App.Session.Pop(r.Context(), "room").(models.Room)
	
	if !ok{
		m.App.ErrorLog.Println("can't get error from session")
		m.App.Session.Put(r.Context(), "error", "can't get room ")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}


	data := make(map[string]interface{})
	data["reservation"]= reservation
	data["room"]= room
	stringMap := make(map[string]string)
	stringMap["startDate"] = reservation.StartDate.Format(DATELAYOUT)
	stringMap["endDate"]   = reservation.EndDate.Format(DATELAYOUT)
    
	render.Template(w, r, "reservation-summary.page.tmpl", models.TemplateData{
	Data:  data,
	StringMap: stringMap ,
	})

 }



 func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	data := models.TemplateData{}
	render.AddDefaultData(&data, r)


    render.Template(w, r, "generals.page.tmpl", data)
	 
	
 }
 func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
 

    render.Template(w, r, "majors.page.tmpl", models.TemplateData{})
	 
	
 }
 func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){

    render.Template(w, r, "contact.page.tmpl", models.TemplateData{})
	 
 }
 
 func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request){
	data := models.TemplateData{}
	render.AddDefaultData(&data, r)
	data.Form =forms.New(nil) 
	
    render.Template(w, r, "search-availability.page.tmpl", data)
 	 
 }

 
 func (m *Repository)PostSearchAvailability(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
 
	if err != nil{
	   helper.ServerError(w, err)
	   return
	}

	form := forms.New(r.PostForm)
	form.TrimValues()

 

	stDate, err := time.Parse(DATELAYOUT, r.PostForm.Get("start_date"))

	if err != nil {
	   form.Errors.Add("start_date", forms.StartDateFormErr)
	}
	
	edDate, err := time.Parse(DATELAYOUT, r.PostForm.Get("end_date"))
	if err != nil {
	   form.Errors.Add("end_date", forms.EndDateFormErr)
	}

 

	

 

	if !form.Valid(){

		render.Template(w, r, "search-availability.page.tmpl", models.TemplateData{
			Form: form,
			CSRFToken: nosurf.Token(r),
		})
		return
	}
 
	
	http.Redirect(w, r, fmt.Sprintf("/choose-room/%s/%s",stDate.Format(DATELAYOUT), edDate.Format(DATELAYOUT)), http.StatusSeeOther)
    
	 
 }

 func (m *Repository)Getrooms(w http.ResponseWriter, r *http.Request){
	data := models.TemplateData{}
	render.AddDefaultData(&data, r)

	pageReq, exists := mux.Vars(r)["page"]
	page := 1.0
	if exists{
		pasedPage, err := strconv.Atoi(pageReq)
		if err == nil && pasedPage > 0{
			page = float64(pasedPage)
		}	 
	}
    
	 countRooms, err := m.DBRepo.CountRooms(r.Context())

	 if err != nil {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/ ", http.StatusTemporaryRedirect)
		return
	}

	pagination := pagination.Pagination{
		 TotalItems: float64(countRooms),
		 PageSize: 10.0,
		 CurrentPage: page,
		 BeforeAndAfterPage: 5,
	}
	totalpageNumber := pagination.TotalPages()
	if totalpageNumber < int(page){
		pagination.CurrentPage = 1.0
	}
	links := pagination.GeneratePageNumbers()

 
	//get rooms pagination

	rooms, err :=  m.DBRepo.GetRoomsByOffset(r.Context(), int(pagination.PageSize), int(pagination.Offset()))

	
    if err != nil || len(rooms)==0 {
		m.App.ErrorLog.Println(err)
		m.App.Session.Put(r.Context(), "error", "there's something wrong happend")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
   data.Data = make(map[string]interface{})
   data.IntMap =make(map[string]int)
   data.Data["rooms"] = rooms
   data.Data["links"] =  links
   data.IntMap["pageNumber"] = totalpageNumber 

   render.Template(w, r, "rooms.page.tmpl", data)
	
 }

 type jsonResponse struct{
	  OK bool `json:"ok"`
	  Message string `json:"message"`
 }
 func (m *Repository)AvailabilityJson(w http.ResponseWriter, r *http.Request){
	
	err := r.ParseForm()
	if err != nil {
	    json.NewEncoder(w).Encode(jsonResponse{OK:false, Message: "there is something wrong with the form please try again later"})
		return
	}
 

	stDate, sterr := time.Parse(DATELAYOUT, r.PostForm.Get("start_date"))

 
	edDate, ederr := time.Parse(DATELAYOUT, r.PostForm.Get("end_date"))


	if sterr != nil || ederr != nil{
 
		json.NewEncoder(w).Encode(jsonResponse{OK:false, Message: "Date not correct"})	
		return
	} 
	roomId, err := strconv.Atoi(r.PostForm.Get("room"))

	if err != nil{
		json.NewEncoder(w).Encode(jsonResponse{OK:false, Message: "wrong room"})	
		return
	} 

	
    notavailable , err := m.DBRepo.CheckAvailabilityByDateAndRoom(r.Context(), roomId, stDate, edDate)

	if err != nil{
		json.NewEncoder(w).Encode(jsonResponse{OK:false, Message: "wrong room"})	
		return
	} 
	 

	if notavailable {
		json.NewEncoder(w).Encode(jsonResponse{OK: false, Message: "Not available"})
		return
	}

	 json.NewEncoder(w).Encode(jsonResponse{OK: true, Message: "Available"})
 }
