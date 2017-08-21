package models

import (
	"github.com/khisakuni/strong/database"
)

type Workout struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (w *Workout) Create() error {
	return database.Conn.Create(&w).Error
}
