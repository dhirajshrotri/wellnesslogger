package routes

import (
	"net/http"

	"github.com/dhirajshrotri/WellnessLogger/handlers"
)

func RegisterExerciseRoutes() {
	http.HandleFunc("/exercise", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateExercise(w, r)
		case http.MethodGet:
			handlers.GetExercises(w, r)
		case http.MethodPut:
			handlers.UpdateExercise(w, r)
		case http.MethodDelete:
			handlers.DeleteExercise(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
