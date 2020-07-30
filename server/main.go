package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controllers"
)

func main() {
	// CRUD Process
	http.HandleFunc("/cretate", controllers.Create)
	http.HandleFunc("/read", controllers.Read)
	http.HandleFunc("/find", controllers.Find)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)

	// Auth Process
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)

	http.ListenAndServe(":8888", nil)
}
