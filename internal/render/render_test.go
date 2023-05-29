package render

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AbdulwahabNour/booking/internal/models"
)




func TestAddDefaultData(t *testing.T){
    var td models.TemplateData
//AddDefaultData(td *models.TemplateData, r *http.Request)  
  r, err := getSession()
  if err != nil{
    t.Error(err)
  }
  session.Put(r.Context(), "flash", "falsh message")

  result :=AddDefaultData(&td, r)
  if result.Flash != "falsh message"{
       t.Errorf("flash message not found in session")
  }
}



func TestTemplate(t *testing.T){
  var td models.TemplateData
    SetTemplatePath("../../templates/")
    cach, err := CreateTemplateCache()
    app.TemplateCache = cach
    if err != nil{
        t.Error(err)
    }
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/", nil)
    err = Template(w , r,"home.page.tmpl",td)
    if err != nil{
       t.Errorf("RenderTemplate returned an error: %v", err)
    }
 
   app.TemplateCache = cach 
 
} 

 

func getSession()(*http.Request, error){
     req, err := http.NewRequest("GET", "/", nil)
   
     if err != nil{
        return nil , err
     }
     ctx := req.Context()
 
     ctx, _ = session.Load(ctx, req.Header.Get("X-Seesion"))
     req = req.WithContext(ctx)
  

    
     return req, nil
}
func TestSetConfigToRender(t *testing.T){
   
   SetConfigToRender(app)
}



