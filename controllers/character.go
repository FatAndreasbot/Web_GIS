package controllers

import (
	"models"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
)

func (s *Server) PostCharacter(c *gin.Context) {
	type newCharacter struct {
		Name      string `json:"Name" binding:"required"`
		MaxHealth int    `json:"MaxHealth" binding:"required"`
		AC        int    `json:"AC" binding:"required"`
		Strenght  int    `json:"Strenght" binding:"required"`
		Dexterity int    `json:"Dexterity" binding:"required"`
	}

	var charData newCharacter
	err := c.BindJSON(&charData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	character := models.Character{
		Name:      charData.Name,
		MaxHealth: charData.MaxHealth,
		Health:    charData.MaxHealth,
		AC:        charData.AC,
		Strenght:  charData.Strenght,
		Dexterity: charData.Dexterity,
		UserID:    user.ID,
	}

	if err := s.db.Create(&character).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": character})
}
