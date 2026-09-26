package utils

import (
	"database/sql"

	"github.com/clinton-mwachia/go-sqlite-mastery/07-update-data/models"
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

func InsertMultiple(db *sql.DB, countries []models.Country) error {
	// Start a transaction for better performance
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Prepare statement once
	stmt, err := tx.Prepare(`INSERT INTO countries (name, population, area) VALUES (?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	// Insert each country
	for _, c := range countries {
		_, err := stmt.Exec(c.Name, c.Population, c.Area)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}

func Update(db *sql.DB, id int, population int) (int64, error) {
	sql := `UPDATE countries SET population = ? WHERE id = ?;`
	result, err := db.Exec(sql, population, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func FindById(db *sql.DB, id int) (*models.Country, error) {
	sql := `SELECT * FROM countries WHERE id = ?`
	row := db.QueryRow(sql, id)
	c := &models.Country{}
	err := row.Scan(&c.Id, &c.Name, &c.Population, &c.Area)
	if err != nil {
		return nil, err
	}
	return c, nil
}
