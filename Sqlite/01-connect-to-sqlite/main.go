package main

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// Connect to the SQLite database
	// this will create the db *test.db* if it doesn't exist
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		// print the error if it exists
		fmt.Println(err)
		return
	}

	// close the connection before the program ends
	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sqliteVersion)
}
