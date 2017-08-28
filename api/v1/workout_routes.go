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
	router.PUT("/api/v1/workouts/:id", workoutUpdateAction)
	router.DELETE("/api/v1/workouts/:id", workoutDestroyAction)
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

func workoutUpdateAction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		InternalServerError(w)
		return
	}

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		InternalServerError(w)
		return
	}
	workout.ID = id
	err = database.Conn.Save(&workout).Error
	if err != nil {
		BadRequestError(w, err.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workout)
}

func workoutDestroyAction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		InternalServerError(w)
		return
	}
	workout := models.Workout{ID: id}
	err = database.Conn.Delete(&workout).Error
	if err != nil {
		BadRequestError(w, err.Error())
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
