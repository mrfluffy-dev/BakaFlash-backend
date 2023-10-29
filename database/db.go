package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func Run(firstName string, lastName string) {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXISTS balls(id INTEGER PRIMARY KEY, firstName TEXT, lastName TEXT)")
	statement.Exec()
	queryStatment, _ := db.Prepare("INSERT INTO balls (firstName, lastName) VALUES (?, ?)")
	queryStatment.Exec(firstName, lastName)
	defer db.Close()
	rows, _ := db.Query("SELECT id, firstName, lastName FROM balls")
	var id int
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName)
		fmt.Println(strconv.Itoa(id) + ": " + firstName + " " + lastName)
	}
}
