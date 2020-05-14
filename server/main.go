package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controller"
)

func main() {
	http.HandleFunc("/", controller.Create)
	http.HandleFunc("/tasks", controller.Read)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/taskdelete", controller.Delete)
	http.ListenAndServe(":8888", nil)
}
