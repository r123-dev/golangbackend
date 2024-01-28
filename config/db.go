package config

import (
	"context"
	"log"

	//"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient is exported Mongo Database client
var MongoDBClient *mongo.Client

// ConnectDatabase is used to connect the MongoDB database
func ConnectDatabase() {
	log.Println("Database connecting...")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://rohanmourya879:vivek123@cluster0.upxigtg.mongodb.net/Vivek")//os.Getenv("DB_URL"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	MongoDBClient = client
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = MongoDBClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected.")
}