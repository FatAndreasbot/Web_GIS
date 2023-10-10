package main

import (
	"controllers"
	"fmt"
	"log"
	"middleware"
	"models"
	"os"
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
	authorized.GET("/update_token", server.UpdateToken)

	authorized.POST("/character", server.PostCharacter)
	authorized.DELETE("/character", server.DeleteCharacter)
	authorized.GET("/character", server.GetCharacter)
	authorized.PATCH("/character", server.UpdateCharacter)

	return r
}

func main() {
	r := SetupRouter()
	utils.SetEnv()

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("DOMAIN_NAME"), "8080"))
}
