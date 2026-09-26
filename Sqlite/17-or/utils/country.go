package utils

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/clinton-mwachia/go-sqlite-mastery/17-or/models"
	_ "github.com/glebarez/go-sqlite"
)

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY,
        name     TEXT NOT NULL,
        population INTEGER NOT NULL,
        area INTEGER NOT NULL
    );`

	return db.Exec(sql)
}

func ReadCSV(filename string) ([]models.Country, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Parse the CSV file
	var countries []models.Country
	for _, record := range records[1:] { // Skip header row
		population, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		area, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		country := models.Country{
			Name:       record[0],
			Population: population,
			Area:       area,
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func Insert(db *sql.DB, c *models.Country) (int64, error) {
	sql := `INSERT INTO countries (name, population, area) 
            VALUES (?, ?, ?);`
	result, err := db.Exec(sql, c.Name, c.Population, c.Area)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func FindByPopulationOrArea(db *sql.DB, minPopulation int, maxArea float64) ([]models.Country, error) {
	query := `SELECT id, name, population, area 
	          FROM countries 
	          WHERE population > ? OR area < ?`

	rows, err := db.Query(query, minPopulation, maxArea)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var c models.Country
		if err := rows.Scan(&c.Id, &c.Name, &c.Population, &c.Area); err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}

	return countries, rows.Err()
}
