package controllers

import (
	"config"
	"models"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	type registerInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input registerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}

	err = user.HashPassword()
	if err != nil {
		c.IndentedJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": user})
}

func Login(c *gin.Context) {

	type loginInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	loginCheck := func(username, password string) (string, error) {
		var err error

		user := models.User{}

		err = config.DB.Model(models.User{}).Where("username = ?", username).Take(&user).Error
		if err != nil {
			return "", err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return "", err
		}

		token, err := utils.GenerateToken(user)
		if err != nil {
			return "", err
		}

		return token, nil
	}

	var input loginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}

	// token, err := loginCheck(user.Username, user.Password)

	token, err := loginCheck(user.Username, user.Password)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "wrong username or password"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
