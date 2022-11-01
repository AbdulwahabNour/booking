package handlers

import (
	"fmt"
	"net/http"

	"github.com/AbdulwahabNour/booking/models"
	"github.com/AbdulwahabNour/booking/pkg/config"
	"github.com/AbdulwahabNour/booking/pkg/render"
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