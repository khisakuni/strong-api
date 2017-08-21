package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var Conn *gorm.DB

func Connect() (*gorm.DB, error) {
	dbsrc := os.Getenv("DATABASE_URL")
	if dbsrc == "" {
		dbsrc = "host=localhost user=koheihisakuni dbname=strong sslmode=disable"
	}
	fmt.Println(dbsrc)

	Conn, err := gorm.Open("postgres", dbsrc)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
