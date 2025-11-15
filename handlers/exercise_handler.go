package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dhirajshrotri/WellnessLogger/config"
	"github.com/dhirajshrotri/WellnessLogger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var exerciseCollection = func() *mongo.Collection {
	// lazy load only when called
	return config.GetCollection("exercises")
}

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	collection := exerciseCollection()

	var exercise models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exercise.Date.IsZero() {
		exercise.Date = time.Now()
	}

	result, err := collection.InsertOne(context.TODO(), exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	// query := bson.M{}
	// if userParam := r.URL.Query().Get("user"); userParam != "" {
	// 	query["UserID"] = bson.M{"eq": userParam}
	// }
	// if dateParam := r.URL.Query().Get("date"); dateParam != "" {
	// 	date, err := time.Parse("2006-01-02", dateParam)
	// 	if err != nil {
	// 		http.Error(w, "invalid date format (YYYY-MM-DD)", http.StatusBadRequest)
	// 		return
	// 	}
	// 	start := date
	// 	end := date.Add(24 * time.Hour)
	// 	query["Date"] = bson.M{"$gte": start, "$lt": end}
	// }

	collection := exerciseCollection()

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var exercises []models.Exercise
	if err = cursor.All(context.TODO(), &exercises); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exercises)
}

func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	collection := exerciseCollection()

	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var updatedData models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	update := bson.M{"$set": updatedData}
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("✅ Updated successfully")
}

func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	collection := exerciseCollection()

	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("🗑️ Deleted successfully")
}
