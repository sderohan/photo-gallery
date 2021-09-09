package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sderohan/photo-gallery/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	//rendering is handled inside the ServeHTTP overriden method inside view.go file
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8001", r)
}
