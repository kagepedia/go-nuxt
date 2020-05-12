package util

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB Start
func Sqlhandler() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(go-nuxt-docker_db_1)/sample_db?parseTime=true")
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
