package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	err := godotenv.Load("/home/rawan/Codescalers/Linktree-RawanMostafa/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := os.Getenv("MONGO_DB_URL")
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions := options.Client().ApplyURI(url)

	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mongoDB!")
	return client

}

var Client *mongo.Client = ConnectToMongo()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
