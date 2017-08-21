package main

import (
	"fmt"
	"github.com/khisakuni/strong/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Errorf("failed to connect to DB: %v\n", err)
	}
	defer db.Close()
}
