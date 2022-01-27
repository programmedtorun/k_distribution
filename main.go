package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

func main() {

	db, err := sql.Open("mysql", "root:Hellodataland123!@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
