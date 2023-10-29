package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func run() {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXIST balls(id INTEGER PRIMARY KEY, firstName TEXT, lastName TEXT)")
	statement.Exec()
	queryStatment, _ := db.Prepare("INSERT INTO balls (firstName, lastName) VALUES (?, ?)")
	queryStatment.Exec("Vet", "Koos")
	defer db.Close()
	rows, _ := db.Query("SELECT id, firstName, lastName FROM balls")
	var id int
	var firstName string
	var lastName string
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName)
		fmt.Println(strconv.Itoa(id) + ": " + firstName + " " + lastName)
	}
}
