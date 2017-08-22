package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Routes sets up v1 routes
func Routes(router *httprouter.Router) {
	WorkoutRoutes(router)
}

type errorMessage struct {
	Message string
	status  int
}

func NotFound(w http.ResponseWriter) {
	err := errorMessage{
		Message: "Resource not found",
		status:  http.StatusNotFound,
	}
	writeErrorJSON(w, err)
}

func InternalServerError(w http.ResponseWriter) {
	err := errorMessage{
		Message: "Something went wrong",
		status:  http.StatusInternalServerError,
	}
	writeErrorJSON(w, err)
}

func BadRequestError(w http.ResponseWriter, message string) {
	err := errorMessage{
		Message: message,
		status:  http.StatusBadRequest,
	}
	writeErrorJSON(w, err)
}

func writeErrorJSON(w http.ResponseWriter, err errorMessage) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(err.status)
	json.NewEncoder(w).Encode(err)
}
