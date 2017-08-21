package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/khisakuni/strong/api/v1"
	"github.com/khisakuni/strong/database"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Errorf("failed to connect to DB: %v\n", err)
	}
	defer db.Close()
	database.Conn = db

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := httprouter.New()
	v1.Routes(router)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
