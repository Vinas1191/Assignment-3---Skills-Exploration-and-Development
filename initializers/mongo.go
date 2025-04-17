package initializers

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func ConnectToMongo() {
    uri := os.Getenv("MONGO_URI")
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Failed to create MongoDB client:", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    MongoClient = client
MongoDatabase = client.Database("Cricket_data") 

log.Println("✅ Connected to MongoDB and selected database: Players")

}
