package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Call cancel to release resources
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	// Connect to the GrubDatabase
	db := client.Database("GrubDatabase")

	log.Println("Successfully connected to the database")
	return db, nil
}

func init() {
	// Initialize MongoDB client  and connect to the database
	_, err := StartDB()
	if err != nil {
		panic(err)
	}
}
