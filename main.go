package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sderohan/photo-gallery/controllers"
	"github.com/sderohan/photo-gallery/models"
)

const (
	host     = "localhost"
	port     = 49154
	user     = "root"
	password = "root"
	dbname   = "photo_gallery"
)

func main() {

	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)
	// Create the connection
	us, err := models.NewUserService(mysqlInfo)
	if err != nil {
		panic(err)
	}
	// Close the connection
	defer us.Close()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	//rendering is handled inside the ServeHTTP overriden method inside view.go file
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8001", r)
}
