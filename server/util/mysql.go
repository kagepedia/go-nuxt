package util

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Sqlhandler() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sample_db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(err)

	return db
}
