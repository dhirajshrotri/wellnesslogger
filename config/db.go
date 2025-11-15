package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectDB() {
	host := getEnv("MONGO_HOST", "mongodb")
	port := getEnv("MONGO_PORT", "27017")
	user := getEnv("MONGO_USER", "")
	pass := getEnv("MONGO_PASS", "")
	dbname := getEnv("MONGO_DB", "mydb")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	DB = client.Database(dbname)
}

func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Database not initialized. Call ConnectDB() first.")
	}
	return DB.Collection(name)
}
