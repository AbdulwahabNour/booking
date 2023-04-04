package forms

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)
 

func TestForm_Require(t *testing.T){

    reqBody := strings.NewReader("first_name=Ahmed&last_name=Mohamed")
  
    r := httptest.NewRequest("POST", "/", reqBody)
    r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    r.ParseForm()
     
    
    form := New(r.PostForm)
 
    form.Require("first_name", "last_name")
    if !form.Valid(){
        for _, v := range form.Errors{
            t.Error(v)
        } 
    }
    v := url.Values{}
    v.Add("first_name","")
    form = New(v)
    form.Require("first_name")

    if form.Valid(){
        t.Error("this form must be not valid")
    }

    

}

func TestForm_minLength(t *testing.T){
    body := strings.NewReader("first_name=AhmedAhmedAhmed&secound_name=Mohamed")

    r := httptest.NewRequest("POST", "/", body)

    r.Header.Add("Content-Type","application/x-www-form-urlencoded")

    r.ParseForm()

    form := New(r.PostForm)
    form.MinLength("first_name", 10)
    if !form.Valid(){
        t.Errorf("this form %s is less than 10 char","first_name" )
    }

    v:= url.Values{}
    v.Add("first_name", "ahmed")
    form = New(v)
    form.MinLength("first_name", 10)
    if form.Valid(){
        t.Errorf("this form %s is less than 10 char and must be not valid","first_name" )
    }



}
func TestForm_IsEmail(t *testing.T){
    body := strings.NewReader("first_name=AhmedAhmedAhmed&Email=Mohamed@yahoo.com")
    r := httptest.NewRequest("POST", "/", body)
    r.Header.Add("Content-Type","application/x-www-form-urlencoded")
    r.ParseForm()
    form := New(r.PostForm)
    form.IsEmail("Email")
    
    if !form.Valid(){
        t.Errorf("%s this is not valid Email", form.Get("Email"))
    }
    
    v := url.Values{}
    v.Add("Email", "ahmedemail")
    form = New(v)
    form.IsEmail("Email")
    if form.Valid(){
        t.Errorf("%s this is not valid Email and it pass", form.Get("Email"))
    }
}

func TestForm_valid(t *testing.T){
    reqBody := strings.NewReader("first_name=Ahmed&secound_name=Mohamed")
    r := httptest.NewRequest("POST","/test", reqBody)


    f := New(r.PostForm)
    if !f.Valid(){
     t.Error("invalid form")
    }
}
func TestForm_TrimValues(t *testing.T){
    valuelist :=map[string][]string{
        "first_name" :[]string{"ahmed"},
        "last_name": []string{"Mohamed"},
        "Email": []string{"ahmedmohamed@yahoo.com"},
    }
    v:= url.Values{}
    v.Add("first_name", " ahmed ")
    v.Add("last_name", " Mohamed ")
    v.Add("Email", "     ahmedmohamed@yahoo.com    ")
    form := New(v)
    form.TrimValues()
    for k, val := range valuelist{
        if v.Get(k) != val[0]{
                t.Errorf("trimValues function didn't trim this '%s'", val[0])
        }
    }
}