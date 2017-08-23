package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/khisakuni/strong/api/v1"
	"github.com/khisakuni/strong/database"
	"github.com/rs/cors"
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

	router.GET("/api/v1/test", Hello)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "HEAD", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
}
