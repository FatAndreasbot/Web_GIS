package main

import (
	"controllers"
	"log"
	"middleware"
	"models"
	"utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DB_Init() *gorm.DB {
	db, err := models.SetupDB()
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

	authorized := r.Group("/user")
	authorized.Use(middleware.JwtAuthMiddleware())
	authorized.POST("/create_character", server.PostCharacter)

	return r
}

func main() {
	r := SetupRouter()
	utils.SetEnv()

	r.Run("localhost:8080")
}
