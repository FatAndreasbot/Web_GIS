package views

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func Player_CRUD_endpoints(r *gin.Engine) {
	r.GET("/players", controllers.GetPlayers)
	r.GET("/players/:id", controllers.GetPlayerById)

	r.POST("/players", controllers.PostPlayers)

	r.PATCH("/players/:id", controllers.PatchPlayerByID)

	r.DELETE("/players/:id", controllers.DeletePlayer)

}
