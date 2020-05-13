package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controller"
)

func main() {
	// util.Sqlhandler()
	http.HandleFunc("/", controller.HandlerIndex)
	http.HandleFunc("/tasks", controller.HandlerTasks)
	http.HandleFunc("/update", controller.HandlerTasksUpdate)
	http.HandleFunc("/taskdelete", controller.HandlerTasksDelete)
	http.ListenAndServe(":8888", nil)
}
