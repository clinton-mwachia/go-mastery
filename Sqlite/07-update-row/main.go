package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/07-update-data/models"
	"github.com/clinton-mwachia/go-sqlite-mastery/07-update-data/utils"
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

	// Update the population of kenya
	_, err = utils.Update(db, 1, 346037975)
	if err != nil {
		fmt.Println(err)
		return
	}
	// confirm updated population
	country, err := utils.FindById(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(country.Population)
}
