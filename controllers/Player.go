package controllers

import (
	"config"
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlayers(c *gin.Context) {
	var players []models.Player
	config.DB.Find(&players)

	c.IndentedJSON(http.StatusOK, gin.H{"data": players})
}

func PostPlayers(c *gin.Context) {
	var input models.Player
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)

	c.IndentedJSON(http.StatusOK, gin.H{"data": input})
}

func GetPlayerById(c *gin.Context) {
	var player models.Player

	err := config.DB.Where("id = ?", c.Param("id")).First(&player).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": player})
}

func PatchPlayerByID(c *gin.Context) {
	var player, input models.Player

	err := config.DB.Where("id = ?", c.Param("id")).First(&player).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	config.DB.Model(&player).Updates(input)

	c.IndentedJSON(http.StatusOK, gin.H{"data": player})
}

func DeletePlayer(c *gin.Context) {
	var player models.Player

	err := config.DB.Where("id = ?", c.Param("id")).First(&player).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Delete(&player)

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
