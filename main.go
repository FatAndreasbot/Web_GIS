package main

import (
	"config"
	"controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.DB_init()

	r.GET("/players", controllers.GetPlayers)
	r.GET("/players/:id", controllers.GetPlayerById)

	r.POST("/players", controllers.PostPlayers)

	r.PATCH("/players/:id", controllers.PatchPlayerByID)

	r.DELETE("/players/:id", controllers.DeletePlayer)

	r.Run("localhost:8080")
}
