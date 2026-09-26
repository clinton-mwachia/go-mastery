package utils

import (
	"database/sql"
	"fmt"

	"github.com/clinton-mwachia/go-sqlite-mastery/30-right-join/models"
	_ "github.com/glebarez/go-sqlite"
)

func CreateTables(db *sql.DB) error {
	countryTable := `
	CREATE TABLE IF NOT EXISTS countries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		population INTEGER NOT NULL,
		area REAL NOT NULL
	);`

	cityTable := `
	CREATE TABLE IF NOT EXISTS cities (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		country_id INTEGER NOT NULL,
		population INTEGER NOT NULL,
		FOREIGN KEY (country_id) REFERENCES countries(id)
	);`

	// Execute both creation statements
	if _, err := db.Exec(countryTable); err != nil {
		return fmt.Errorf("failed to create countries table: %w", err)
	}
	if _, err := db.Exec(cityTable); err != nil {
		return fmt.Errorf("failed to create cities table: %w", err)
	}

	return nil
}

func InsertCountry(db *sql.DB, c *models.Country) (int64, error) {
	sql := `INSERT INTO countries (name, population, area) 
            VALUES (?, ?, ?);`
	result, err := db.Exec(sql, c.Name, c.Population, c.Area)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func InsertCity(db *sql.DB, name string, countryID int, population int) (int64, error) {
	query := `INSERT INTO cities (name, country_id, population) VALUES (?, ?, ?);`
	result, err := db.Exec(query, name, countryID, population)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func InsertSampleData(db *sql.DB) error {
	// Sample countries
	countries := []models.Country{
		{Name: "Kenya", Population: 55000000, Area: 580367},
		{Name: "Tanzania", Population: 63000000, Area: 945087},
		{Name: "Uganda", Population: 48000000, Area: 241038},
	}

	for _, c := range countries {
		id, err := InsertCountry(db, &c)
		if err != nil {
			return fmt.Errorf("failed to insert country %s: %w", c.Name, err)
		}

		// Add sample cities for each country
		switch c.Name {
		case "Kenya":
			InsertCity(db, "Nairobi", int(id), 4500000)
			InsertCity(db, "Mombasa", int(id), 1200000)
		case "Tanzania":
			InsertCity(db, "Dar es Salaam", int(id), 6000000)
			InsertCity(db, "Arusha", int(id), 416000)
		case "Uganda":
			InsertCity(db, "Kampala", int(id), 1600000)
		}
	}

	return nil
}

// A RIGHT JOIN returns all rows from the right table (countries) and matching rows from
// the left table (cities).
// If a country has no cities, then the city fields will be NULL, so we must scan
// using sql.NullString or sql.NullInt64.
// Right Join returns all countries, even those without matching cities
func GetCountriesWithOptionalCities(db *sql.DB) ([]models.City, error) {
	query := `
		SELECT c.name AS country_name, ci.name AS city_name, ci.population
		FROM cities AS ci
		RIGHT JOIN countries AS c ON ci.country_id = c.id
		ORDER BY c.name
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.City

	for rows.Next() {
		var (
			item     models.City
			nullCity sql.NullString
			nullPop  sql.NullInt64
		)

		// city fields may be NULL if no match
		if err := rows.Scan(&item.CountryName, &nullCity, &nullPop); err != nil {
			return nil, err
		}

		if nullCity.Valid {
			item.CityName = nullCity.String
		} else {
			item.CityName = "No City"
		}

		if nullPop.Valid {
			item.CityPop = int(nullPop.Int64)
		} else {
			item.CityPop = 0
		}

		result = append(result, item)
	}

	return result, rows.Err()
}
