package main

import (
	"fmt"
	"net/http"

	"github.com/dhirajshrotri/WellnessLogger/config"
	"github.com/dhirajshrotri/WellnessLogger/routes"
)

func main() {
	config.ConnectDB()          // ✅ Initialize Mongo before using it
	routes.RegisterExerciseRoutes()

	fmt.Println("🚀 Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
