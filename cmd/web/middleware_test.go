package main

import (
	"net/http"
	"testing"
)




func TestNoSurf(t *testing.T){
  var myH testHandler

    returnHandler := NoSurf(&myH)

    switch typehandler := returnHandler.(type){
    case  http.Handler:
         //do nothing
    default:
        t.Errorf("%t Is not type handler", typehandler)
    }
}