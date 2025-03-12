## Database Interactions

To interact with an SQLite database in Go, you can utilize the `database/sql` package along with the `github.com/mattn/go-sqlite3` driver. Below is a comprehensive guide to setting up and performing basic database operations.

**1. Setup and Dependencies**

First, ensure you have Go installed. Initialize a new Go module for your project:

```go
go mod init crud
```

Next, install the SQLite driver:

```go
go get github.com/mattn/go-sqlite3
```

**2. Import Packages**

Import the necessary packages in your Go file:

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)
```

**3. Initialize Database and Create Table**

Establish a connection to the SQLite database and create a table:

```go
func initializeDatabase() (*sql.DB, error) {
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
```

**4. CRUD Operations**

- **Create**

  Insert a new record into the `users` table:

  ```go
  func createUser(db *sql.DB, name string, email string) error {
      insertUserSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
      _, err := db.Exec(insertUserSQL, name, email)
      return err
  }
  ```

- **Read**

  Retrieve all records from the `users` table:

  ```go
  type User struct {
      ID    int
      Name  string
      Email string
  }

  func getUsers(db *sql.DB) ([]User, error) {
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
  ```

- **Update**

  Modify an existing record in the `users` table:

  ```go
  func updateUser(db *sql.DB, id int, name string, email string) error {
      updateUserSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
      _, err := db.Exec(updateUserSQL, name, email, id)
      return err
  }
  ```

- **Delete**

  Remove a record from the `users` table:

  ```go
  func deleteUser(db *sql.DB, id int) error {
      deleteUserSQL := `DELETE FROM users WHERE id = ?`
      _, err := db.Exec(deleteUserSQL, id)
      return err
  }
  ```

**5. Main Function**

Integrate the CRUD functions within the `main` function:

```go
func main() {
    db, err := initializeDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a new user
    err = createUser(db, "Alice", "alice@example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Read users
    users, err := getUsers(db)
    if err != nil {
        log.Fatal(err)
    }
    for _, user := range users {
        fmt.Printf("%d: %s (%s)\n", user.ID, user.Name, user.Email)
    }

    // Update a user
    err = updateUser(db, 1, "Alice Smith", "alice.smith@example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Delete a user
    err = deleteUser(db, 1)
    if err != nil {
        log.Fatal(err)
    }
}
```
