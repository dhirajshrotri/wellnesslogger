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

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
