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

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func GetUsers() []Person {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXISTS balls(id INTEGER PRIMARY KEY, firstName TEXT, lastName TEXT)")
	statement.Exec()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM balls")
	personSlice := []Person{}
	var id int
	var firstName string
	var lastName string
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName)
		var dbUser Person
		dbUser.Id = id
		dbUser.FirstName = firstName
		dbUser.LastName = lastName
		personSlice = append(personSlice, dbUser)
	}
	return personSlice
}

// upload image to db
func UploadImage(imageName string, imageType string, image []byte) {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXISTS images(id INTEGER PRIMARY KEY, imageName TEXT, imageType TEXT, image BLOB)")
	statement.Exec()
	queryStatment, _ := db.Prepare("INSERT INTO images (imageName, imageType, image) VALUES (?, ?, ?)")
	queryStatment.Exec(imageName, imageType, image)
	defer db.Close()
}

func GetImage(imageName string) []byte {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXISTS images(id INTEGER PRIMARY KEY, imageName TEXT, imageType TEXT, image BLOB)")
	statement.Exec()
	defer db.Close()
	rows, _ := db.Query("SELECT image FROM images WHERE imageName = ?", imageName)
	var image []byte
	for rows.Next() {
		rows.Scan(&image)
	}
	return image
}

func ListImageNames() []string {
	db, _ := sql.Open("sqlite3", "./srs.db")
	statement, _ := db.Prepare("CREATE TABLE if NOT EXISTS images(id INTEGER PRIMARY KEY, imageName TEXT, imageType TEXT, image BLOB)")
	statement.Exec()
	defer db.Close()
	rows, _ := db.Query("SELECT imageName FROM images")
	var imageNames []string
	var imageName string
	for rows.Next() {
		rows.Scan(&imageName)
		imageNames = append(imageNames, imageName)
	}
	return imageNames
}
