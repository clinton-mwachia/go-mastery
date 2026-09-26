package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/27-exists/utils"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// connect to the SQLite database
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// create the countries table
	_, err = utils.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table countries was created successfully.")
	// read the CSV file
	countries, err := utils.ReadCSV("countries.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert the data into the SQLite database
	for _, country := range countries {
		_, err := utils.Insert(db, &country)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	exists, err := utils.CountryExists(db, "Kenya")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if exists {
		fmt.Println("Country exists in the database.")
	} else {
		fmt.Println("Country not found.")
	}

}
