package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUsers() {
	fmt.Println("START")
	db, err := sql.Open("sqlite3", "./srs.db")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE if NOT EXISTS balls(id INTEGER PRIMARY KEY, firstName TEXT, lastName TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM balls")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var people []Person
	var id int
	var firstName string
	var lastName string

	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{ID: id, FirstName: firstName, LastName: lastName})
	}

	fmt.Println("END")

	// Printing data from JSON
	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	var peopleFromJSON []Person
	err = json.Unmarshal(jsonData, &peopleFromJSON)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range peopleFromJSON {
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s\n", p.ID, p.FirstName, p.LastName)
	}
}
