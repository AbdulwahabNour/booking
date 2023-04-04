package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)



 func TestMain(m *testing.M){
	
	os.Exit(m.Run())
 }



type testHandler struct{}

func(th *testHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){

	 fmt.Fprint(w, "done")
}