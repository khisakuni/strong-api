package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/khisakuni/strong/database"
	"github.com/khisakuni/strong/models"
	"net/http"
	"strconv"
)

func WorkoutRoutes(router *httprouter.Router) {
	router.GET("/api/v1/workouts", workoutIndexAction)
	router.GET("/api/v1/workouts/:id", workoutShowAction)
	router.POST("/api/v1/workouts", workoutCreateAction)
}

func workoutIndexAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var workouts []models.Workout
	database.Conn.Find(&workouts)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(workouts)
}

func workoutShowAction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		InternalServerError(w)
		return
	}

	workout := models.Workout{ID: id}
	err = database.Conn.Find(&workout).Error
	if err != nil {
		NotFound(w)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(workout)
}

func workoutCreateAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		InternalServerError(w)
		return
	}

	if workout.Name == "" {
		BadRequestError(w, "Name can't be blank")
		return
	}

	if workout.Description == "" {
		BadRequestError(w, "Description can't be blank")
		return
	}

	err := workout.Create()

	if err != nil {
		InternalServerError(w)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workout)
}
