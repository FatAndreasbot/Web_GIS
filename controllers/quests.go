package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetQuests(c *gin.Context) {
	var quests []models.Quest

	if err := s.db.Find(&quests).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quests)
}

func (s *Server) PostQuest(c *gin.Context) {
	type new_Quest struct {
		Name       string `json:"Name" binding:"required"`
		Difficulty int    `json:"Difficulty" binding:"required"`
	}

	var quest new_Quest

	if err := c.ShouldBindJSON(&quest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQuest := models.Quest{Name: quest.Name, Difficulty: quest.Difficulty}

	if err := s.db.Create(&newQuest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newQuest)
}
