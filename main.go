package main

import (
	"github.com/Vinas1191/Assignment-3---Skills-Exploration-and-Development/controllers"
	"github.com/Vinas1191/Assignment-3---Skills-Exploration-and-Development/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	  initializers.LoadEnvVariables()
	  initializers.ConnectToMongo()
}

func main() {
    router := gin.Default()

    router.GET("/players", controllers.GetPlayers)
	router.GET("/players/:id", controllers.GetPlayerByID)
	router.POST("/players", controllers.CreatePlayer)
	router.PUT("/players/:id", controllers.UpdatePlayer)
	router.DELETE("/players/:id", controllers.DeletePlayer) 

    router.Run()
}