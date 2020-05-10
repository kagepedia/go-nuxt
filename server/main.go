package main

import (
	"net/http"

	"github.com/kagepedia/go-api/controller"
	"github.com/kagepedia/go-api/util"
)

func main() {
	util.Sqlhandler()
	http.HandleFunc("/tasks", controller.HandlerTasks)
	http.ListenAndServe(":8888", nil)
}
