package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)
type postData struct{
    key string
    value string
}

var testhand=[]struct{
    name string
    url string
    method string
    params []postData
    expectedStatusCode int
}{
    {"home", "/", "GET", []postData{}, http.StatusOK},
    {"about", "/about", "GET", []postData{}, http.StatusOK},
    {"Generals", "/generals-quarters", "GET", []postData{}, http.StatusOK},
    {"Majors", "/majors-suite", "GET", []postData{}, http.StatusOK},
    {"SearchAvailability", "/search-availability", "GET", []postData{}, http.StatusOK},
    {"Reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
    {"Contact", "/contact", "GET", []postData{}, http.StatusOK},
    {"ReservationSummary", "/reservation-summary","GET",[]postData{}, http.StatusOK},

 
    {"PostSearchAvailability", "/search-availability", "POST", []postData{
        {key: "start", value: "2020-01-01"},
        {key: "end", value: "2020-01-02"}},
         http.StatusOK},

    {"AvailabilityJson", "/search-availability-json", "POST", []postData{
        {key: "start", value: "2020-01-01"},
        {key: "end", value: "2020-01-02"},
    }, http.StatusOK},
    
    {"PostReservation", "/make-reservation", "POST", []postData{
        {key: "first_name", value: "ahmed"},
        {key: "last_name", value: "Mohamed"},
        {key: "email", value: "ahmedmohamed@yahoo.com"},
        {key: "phone", value: "011500013000"},
    }, http.StatusOK},

   
}

func TestHandler(t *testing.T){
 h := getRoutes()
 ts := httptest.NewTLSServer(h)
 
 defer ts.Close()
 for _, v := range testhand{
    if v.method == "GET"{
        res, err  := ts.Client().Get(ts.URL + v.url)
        if err != nil{
            t.Error(err)
            log.Fatal(err)
        }
        if res.StatusCode != v.expectedStatusCode{
            t.Errorf("for %s , expected %d but got %d", v.name, v.expectedStatusCode, res.StatusCode)
        }
    }

  if v.method == "POST"{
    values := url.Values{}
    for _, x:= range v.params{
        values.Add(x.key, x.value)
    }

    res, err := ts.Client().PostForm(ts.URL+v.url, values)
 
    if err != nil{
        t.Error(err)
        log.Fatal(err)
    }

    if res.StatusCode != v.expectedStatusCode{
         t.Errorf("for %s, expected %d but got %d", v.name, v.expectedStatusCode, res.StatusCode)
    }

  }
   


 }
}