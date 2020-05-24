package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controllers"
)

func main() {
	http.HandleFunc("/cretate", controllers.Create)
	http.HandleFunc("/read", controllers.Read)
	http.HandleFunc("/find", controllers.Find)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8888", nil)
}
