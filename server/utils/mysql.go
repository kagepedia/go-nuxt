package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB Start
func Sqlhandler() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(go-nuxt-docker_db_1)/sample_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(err)

	return db
}

// SELECT

// INSERT

// UPDATE

// DELETE
