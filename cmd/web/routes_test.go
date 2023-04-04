package main

import (
	"testing"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/gorilla/mux"
)



func TestRoutes(t *testing.T){

    conf := config.AppConfig{}
    returnHandler := routes(&conf)
    switch v := returnHandler.(type){
    case *mux.Router:
        //do nothing
    default:
        t.Errorf("routes function must return *mux.Router bt it return %T", v)
    }
}

