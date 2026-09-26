package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	idString := "6513470af65dc3213a7e2311"

	objID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		log.Fatalf("Invalid hex string: %v", err)
	}

	fmt.Printf("Successfully converted! Type: %T, Value: %v\n", objID, objID)
}
