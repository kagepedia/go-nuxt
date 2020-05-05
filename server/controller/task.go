package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kagepedia/go-api/util"
)

type Ping struct {
	Status int
	Rssult string
}

type Task struct {
	Id       int64
	Task     string
	CreateAt time.Time
	UpdateAt time.Time
}

var tasks []Task

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "bad"}
	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// 構造体に関して勉強
func HandlerTasks(w http.ResponseWriter, r *http.Request) {
	// json定義
	w.Header().Set("Content-Type", "application/json")

	db := util.Sqlhandler()
	rows, err := db.Query("SELECT id, task FROM t_task WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Task)

		res, err := json.Marshal(task)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(res))

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
