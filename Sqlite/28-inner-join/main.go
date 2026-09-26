package main

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/28-inner-join/utils"
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
	err = utils.CreateTables(db)
	if err != nil {
		fmt.Println("Error creating tables:", err)
	} else {
		fmt.Println("Tables created successfully!")
	}
	err = utils.InsertSampleData(db)
	if err != nil {
		fmt.Println("Error inserting sample data:", err)
		return
	}

	fmt.Println("Sample data inserted successfully!")

	records, err := utils.GetCitiesWithCountries(db)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, r := range records {
		fmt.Printf("Country: %s | City: %s | City Population: %d\n",
			r.CountryName, r.CityName, r.CityPop)
	}

}
