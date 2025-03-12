package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int
	Name  string
	Email string
}

// initialize an sqlite database
func InitializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "email" TEXT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// create a new user
func CreateUser(db *sql.DB, name string, email string) error {
	insertUserSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := db.Exec(insertUserSQL, name, email)
	return err
}

// list all users
func GetUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// update user details
func UpdateUser(db *sql.DB, id int, name string, email string) error {
	updateUserSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := db.Exec(updateUserSQL, name, email, id)
	return err
}

// delete user
func DeleteUser(db *sql.DB, id int) error {
	deleteUserSQL := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(deleteUserSQL, id)
	return err
}
