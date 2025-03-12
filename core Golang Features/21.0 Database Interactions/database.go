package main

import (
	"crud/utils"
	"fmt"
	"log"
)

func main() {
	db, err := utils.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new user
	err = utils.CreateUser(db, "Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Read users
	users, err := utils.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("%d: %s (%s)\n", user.ID, user.Name, user.Email)
	}

	// Update a user
	err = utils.UpdateUser(db, 1, "Alice Smith", "alice.smith@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Delete a user
	err = utils.DeleteUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
}
