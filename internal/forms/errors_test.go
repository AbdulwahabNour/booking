package forms

import "testing"


func TestError_Add(t *testing.T){
     err := make(errors)
    err.Add("first_name", "Ahmed")

    _, ok := err["first_name"]
    if !ok{
        t.Error("first_name field not found", )
    }
}
func TestError_Get(t *testing.T){
     err:= make(errors)
    err["first_name"] =[]string{"firs name must be more than 10 char"}

    
    if (err.Get("first_name")==""){
        t.Error("first_name field not found", )
    }
}