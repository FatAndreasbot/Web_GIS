package controllers

import (
	"models"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Register(c *gin.Context) {
	type registerInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input registerInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{Username: input.Username, Password: input.Password}
	user.HashPassword()

	err = s.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "data": user})
}

func (s *Server) Login(c *gin.Context) {
	type loginInput struct {
		Username string `json:"Username" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	loginCheck := func(username, password string) (string, error) {
		var err error

		user := models.User{}

		if err = s.db.Model(models.User{}).Where("username=?", username).Take(&user).Error; err != nil {
			return "", err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return "", err
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			return "", err
		}

		return token, nil
	}

	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}

	token, err := loginCheck(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The username or password is not correct"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
