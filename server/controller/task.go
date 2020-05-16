/*
** Task CRUD
 */

package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kagepedia/go-api/util"
)

type Task struct {
	Id       int64
	Task     string
	CreateAt time.Time
	UpdateAt time.Time
}

// sample
func Sample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nowTime := time.Now()
	const format = "2006/01/02 15:04:05"
	nowTime.Format(format)
	sample := make(map[string]interface{})
	sample["lang"] = "Go"
	sample["id"] = 9999
	sample["job"] = "professional baseball player"
	sample["time"] = nowTime.Format(format)
	json.NewEncoder(w).Encode(sample)
}

// 構造体に関して勉強（作成）Create まだ
func Create(w http.ResponseWriter, r *http.Request) {

	task := r.FormValue("task")
	nowTime := time.Now()
	const format = "2006/01/02 15:04:05"
	nowTime.Format(format)
	db := util.Sqlhandler()
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO t_task (task,created_at,updated_at) VALUES (?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ins.Close()
	_, err = ins.Exec(&task, nowTime, nowTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// 構造体に関して勉強（表示）Read OK
func Read(w http.ResponseWriter, r *http.Request) {
	// json定義
	w.Header().Set("Content-Type", "application/json")

	db := util.Sqlhandler()
	rows, err := db.Query("SELECT * FROM t_task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Task, &task.CreateAt, &task.UpdateAt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tasks)

	defer db.Close()
}

// 構造体に関して勉強（作成）Find　まだ
func Find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := util.Sqlhandler()
	id := r.FormValue("id")
	var task Task
	err := db.QueryRow("SELECT * FROM t_task WHERE id = ?", id).Scan(&task.Id, &task.Task, &task.CreateAt, &task.UpdateAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// 構造体に関して勉強（作成）Update　まだ
func Update(w http.ResponseWriter, r *http.Request) {
	//update
	nowTime := time.Now()
	db := util.Sqlhandler()
	id := r.FormValue("id")
	task := r.FormValue("task")
	defer db.Close()
	upd, err := db.Prepare("UPDATE t_task SET task = ?, updated_at = ? WHERE id = ? ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer upd.Close()
	upd.Exec(task, nowTime, id)
}

// 構造体に関して勉強（作成）Delete まだ
func Delete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := util.Sqlhandler()
	id := r.FormValue("id")
	deleteData, err := db.Prepare("DELETE FROM t_task WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer deleteData.Close()

	result, err := deleteData.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)

	defer db.Close()

	json.NewEncoder(w).Encode(rowsAffect)
}
