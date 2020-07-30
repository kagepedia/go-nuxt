/*
** User Auth
 */

package controllers

import (
	"encoding/json"
	"github.com/kagepedia/go-api/models"
	"net/http"

	"github.com/kagepedia/go-api/utils"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {

	db := utils.Sqlhandler()
	ID := r.FormValue("ID")
	PW := r.FormValue("PW")
	user := models.User{}
	// var user User
	err := db.QueryRow("SELECT pk, ID FROM t_user WHERE ID = ? and PW = ?", ID, PW).Scan(&user.Pk, &user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Logout
func Logout(w http.ResponseWriter, r *http.Request) {

	db := utils.Sqlhandler()
	ID := r.FormValue("ID")
	PW := r.FormValue("PW")
	user := models.User{}
	// var user User
	err := db.QueryRow("SELECT * FROM t_user WHERE ID = ? and PW = ?", ID, PW).Scan()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
