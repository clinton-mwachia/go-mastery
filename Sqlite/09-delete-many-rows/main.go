package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/09-delete-many-rows/models"
	"github.com/clinton-mwachia/go-sqlite-mastery/09-delete-many-rows/utils"
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

	// create a countries
	countries := []models.Country{
		{Name: "Kenya", Population: 329064917, Area: 9826675},
		{Name: "uganda", Population: 229064917, Area: 1826675},
		{Name: "Tanzania", Population: 129064917, Area: 9006675},
	}

	// insert the countries
	err = utils.InsertMultiple(db, countries)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print the inserted country
	fmt.Println("Multiple countries inserted")

	ids := []int{1, 2}

	// Delete the country with id 1
	rows, err := utils.DeleteMany(db, ids)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d rows deleted\n", rows)
	// confirm deleted countries
	country, err := utils.FindById(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(country.Name)
}
