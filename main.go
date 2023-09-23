package main

import (
	"endpoints"
	"orm"

	"github.com/gin-gonic/gin"
)

func main() {

	db := orm.SqliteConnection{ConnString: "DB.db"}
	db.Connect()

	router := gin.Default()
	router.GET("/players", endpoints.GetPlayers)

	router.Run("localhost:8080")
}
