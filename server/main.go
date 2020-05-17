package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controller"
)

func main() {
	http.HandleFunc("/sample", controller.Sample)
	http.HandleFunc("/cretate", controller.Create)
	http.HandleFunc("/read", controller.Read)
	http.HandleFunc("/readsample", controller.ReadSample)
	http.HandleFunc("/find", controller.Find)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)
	http.ListenAndServe(":8888", nil)
}
