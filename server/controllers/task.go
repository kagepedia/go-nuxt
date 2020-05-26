/*
** Task CRUD
 */

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kagepedia/go-api/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kagepedia/go-api/utils"
)

// Create
func Create(w http.ResponseWriter, r *http.Request) {

	task := models.Task{}
	task.Task = r.FormValue("task")
	nowTime := time.Now()
	const format = "2006/01/02 15:04:05"
	nowTime.Format(format)
	db := utils.Sqlhandler()
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO t_task (task,created_at,updated_at) VALUES (?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ins.Close()
	_, err = ins.Exec(&task.Task, nowTime, nowTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Read
func Read(w http.ResponseWriter, r *http.Request) {

	db := utils.Sqlhandler()
	rows, err := db.Query("SELECT * FROM t_task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

	defer db.Close()
}

// Find
func Find(w http.ResponseWriter, r *http.Request) {

	db := utils.Sqlhandler()
	id := r.FormValue("id")
	task := models.Task{}
	// var task Task
	err := db.QueryRow("SELECT * FROM t_task WHERE id = ?", id).Scan(&task.Id, &task.Task, &task.CreateAt, &task.UpdateAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Update
func Update(w http.ResponseWriter, r *http.Request) {
	//update
	nowTime := time.Now()
	db := utils.Sqlhandler()
	task := models.Task{}
	task.Id, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	task.Task = r.FormValue("task")
	defer db.Close()
	upd, err := db.Prepare("UPDATE t_task SET task = ?, updated_at = ? WHERE id = ? ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer upd.Close()
	upd.Exec(&task.Task, nowTime, &task.Id)
}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {

	db := utils.Sqlhandler()
	task := models.Task{}
	task.Id, _ = strconv.ParseInt(r.FormValue("id"), 10, 64)
	deleteData, err := db.Prepare("DELETE FROM t_task WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer deleteData.Close()

	result, err := deleteData.Exec(&task.Id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffect)
}
