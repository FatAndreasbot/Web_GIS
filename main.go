package main

import (
	"controllers"
	"log"
	"utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DB_Init() *gorm.DB {
	db, err := utils.SetupDB()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := DB_Init()
	server := controllers.NewServer(db)

	r.POST("/register", server.Register)
	r.POST("/login", server.Login)

	return r
}

func main() {
	r := SetupRouter()
	utils.SetEnv()

	r.Run("localhost:8080")
}
