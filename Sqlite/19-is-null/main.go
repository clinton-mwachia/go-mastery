package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/19-is-null/utils"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// connect to the SQLite database
	db, err := sql.Open("sqlite", "./test.db?_pragma=foreign_keys(1)")
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

	countries, err = utils.FindWhereAreaIsNull(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(countries) == 0 {
		fmt.Println("No countries found in that range")
	}

	for _, c := range countries {
		fmt.Printf("%s-%d\n", c.Name, c.Population)
	}
}
