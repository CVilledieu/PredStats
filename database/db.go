package database

import (
	"database/sql"
	"log"
)

const (
	connectString = "user=postgres password=postgres dbname=postgres sslmode=disable"
)

func DBConn() error {
	// This function is used to use the database
	db, err := sql.Open("postgres", connectString)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return nil
}

func postData() {
	// This function is used to import data from the API into the database
}

func getData() {
	// This function is used to get data from the database
}

func updateData() {
	// This function is used to update data in the database
}

func deleteData() {
	// This function is used to delete data from the database
}
