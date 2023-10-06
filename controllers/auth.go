package controllers

import (
	"models"
	"net/http"
	"os"
	"strconv"
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
		Login    string `json:"login" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input registerInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{Username: input.Username, Password: input.Password, Login: input.Login}
	user.HashPassword()

	err = s.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "data": user})
}

func (s *Server) Login(c *gin.Context) {
	type loginInput struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	loginCheck := func(login, password string) (string, error) {
		var err error

		user := models.User{}

		if err = s.db.Model(models.User{}).Where("login=?", login).Take(&user).Error; err != nil {
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

	user := models.User{Login: input.Login, Password: input.Password}

	token, err := loginCheck(user.Login, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The username or password is not correct"})
		return
	}

	// send cookie to user
	// c.SetCookie("auth_cookie", token, 604800, "/", "localhost", true, true)

	tokenHourLifespan, _ := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	c.SetCookie("gin_auth_cookie", token,
		3600*tokenHourLifespan, "/", os.Getenv("DOMAIN_NAME"), true, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) UpdateToken(c *gin.Context) {
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenHourLifespan, _ := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	c.SetCookie("gin_auth_cookie", token,
		3600*tokenHourLifespan, "/", os.Getenv("DOMAIN_NAME"), true, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
