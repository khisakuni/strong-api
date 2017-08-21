package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/khisakuni/strong/models"
	"net/http"
)

func WorkoutRoutes(router *httprouter.Router) {
	//router.GET("/api/v1/workouts", workoutsIndex)
	//router.GET("/api/v1/workouts/:id", workoutsShow)
	router.POST("/api/v1/workouts", workoutCreateAction)
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
