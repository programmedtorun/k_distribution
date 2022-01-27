package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DThing struct {
	DeptNo   string `json:"id"`
	DeptName string `json:"name"`
}

func main() {
	fmt.Println("What is the mysqlDB password?") // for password, reference Patrick's mysql "LastPass" app
	var ipt string
	fmt.Println("\033[8m")
	fmt.Scanln(&ipt)
	fmt.Println("\033[28m")
	eDBAccess := "root:" + ipt + "@tcp(127.0.0.1:3306)/employees"
	db, err := sql.Open("mysql", eDBAccess)
	if err != nil {
		fmt.Println("There was an error opening the sql db:  ", err)
	}

	defer db.Close()
	// Execute the query - this is a test query that has some random test data I've added to mysql on my comp
	results, err := db.Query("SELECT * FROM employees.departments")
	if err != nil {
		fmt.Println("There was an error with the sql query:  ", err)
	}

	for results.Next() {
		var dThing DThing
		// for each row, scan the result into composite object
		err = results.Scan(&dThing.DeptNo, &dThing.DeptName)
		if err != nil {
			fmt.Println("There was an error scanning data into DThing struct:  ", err)
		}
		// and then print out the DeptName attribute
		fmt.Println(dThing.DeptName)
	}
}
