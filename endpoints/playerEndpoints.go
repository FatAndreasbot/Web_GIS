package endpoints

import (
	"models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var p = models.Player{
	Name:      "Andy",
	ModelData: gorm.Model{ID: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetPlayers(c *gin.Context) {
	// players := []map[string]any{{"Id":p.Id, "Name": p.Name}}
	players := []models.Player{p}
	c.IndentedJSON(http.StatusOK, players)
}
