package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request){
var filename = "login.html"
 t, err := template.ParseFiles(filename)
 if err != nil {
	fmt.Println("error while parsing files", err)
	return
 }
  err =t.ExecuteTemplate(w , filename, nil)
  if err != nil {
	fmt.Println("error while executing tamplate" , err)
	return
  }
}

func loginsubmit(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username , password)
}

func handler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/login":
		login(w,r)
    case "/login-submit":
    	loginsubmit(w,r)
    default:
    	fmt.Fprintf(w,"sup ninjas")
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("" , nil)
}