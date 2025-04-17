package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Vinas1191/Assignment-3---Skills-Exploration-and-Development/initializers"
	"github.com/Vinas1191/Assignment-3---Skills-Exploration-and-Development/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPlayers(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := initializers.MongoDatabase.Collection("players")

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching players"})
        return
    }
    defer cursor.Close(ctx)

    var players []bson.M
    if err := cursor.All(ctx, &players); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding players"})
        return
    }

    c.JSON(http.StatusOK, players)
}

func GetPlayerByID(c *gin.Context) {
	playerID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		log.Println("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	collection := initializers.MongoDatabase.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		log.Println("Player not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreatePlayer(c *gin.Context) {
	var newPlayer struct {
		Age         int       `json:"age"`
		Batting     string    `json:"batting"`
		Bowling     string    `json:"bowling"`
		Centuries   int       `json:"centuries"`
		DateOfBirth string    `json:"dateOfBirth"`
		HatTricks   int       `json:"hatTricks"`
		JerseyNumber int      `json:"jerseyNumber"`
		Name        string    `json:"name"`
		Role        string    `json:"role"`
		Runs        int       `json:"runs"`
		TeamsPlayedFor []struct {
			Matches int    `json:"matches"`
			Team    string `json:"team"`
		} `json:"teamsPlayedFor"`
		Wickets int `json:"wickets"`
	}

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := initializers.MongoDatabase.Collection("players")

	player := bson.D{
		bson.E{Key: "age", Value: newPlayer.Age},
		bson.E{Key: "batting", Value: newPlayer.Batting},
		bson.E{Key: "bowling", Value: newPlayer.Bowling},
		bson.E{Key: "centuries", Value: newPlayer.Centuries},
		bson.E{Key: "dateOfBirth", Value: newPlayer.DateOfBirth},
		bson.E{Key: "hatTricks", Value: newPlayer.HatTricks},
		bson.E{Key: "jerseyNumber", Value: newPlayer.JerseyNumber},
		bson.E{Key: "name", Value: newPlayer.Name},
		bson.E{Key: "role", Value: newPlayer.Role},
		bson.E{Key: "runs", Value: newPlayer.Runs},
		bson.E{Key: "teamsPlayedFor", Value: newPlayer.TeamsPlayedFor},
		bson.E{Key: "wickets", Value: newPlayer.Wickets},
	}

	_, err := collection.InsertOne(c, player)
	if err != nil {
		log.Fatal("Error inserting player:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully!"})
}

func UpdatePlayer(c *gin.Context) {
    playerID := c.Param("id")

    objectID, err := primitive.ObjectIDFromHex(playerID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
        return
    }

    var updatedPlayer models.Player
    if err := c.ShouldBindJSON(&updatedPlayer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	collection := initializers.MongoDatabase.Collection("players")

    filter := bson.M{"_id": objectID}

    update := bson.M{
        "$set": bson.M{
            "age":          updatedPlayer.Age,
            "batting":      updatedPlayer.Batting,
            "bowling":      updatedPlayer.Bowling,
            "centuries":    updatedPlayer.Centuries,
            "dateOfBirth":  updatedPlayer.DateOfBirth,
            "hatTricks":    updatedPlayer.HatTricks,
            "jerseyNumber": updatedPlayer.JerseyNumber,
            "name":         updatedPlayer.Name,
            "role":         updatedPlayer.Role,
            "runs":         updatedPlayer.Runs,
            "teamsPlayedFor": updatedPlayer.TeamsPlayedFor,
            "wickets":      updatedPlayer.Wickets,
        },
    }

    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
        return
    }

    if result.MatchedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Player not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Player updated successfully"})
}

func DeletePlayer(c *gin.Context) {
	playerID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		log.Println("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	collection := initializers.MongoDatabase.Collection("players")
	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		log.Println("Error deleting player:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}