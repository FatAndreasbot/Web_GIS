package utils

import (
	"errors"
	"fmt"
	"models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {

	tokenLifespanEnvVar := os.Getenv("TOKEN_HOUR_LIFESPAN")
	tokenLifespan, err := strconv.Atoi(tokenLifespanEnvVar)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SERCET")))

}

func ValidateToken(c *gin.Context) error {
	token, err := GetToken(c)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token provided")
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	getTokenFromRequest := func(c *gin.Context) string {
		bearerToken := c.Request.Header.Get("Authorization")

		splitToken := strings.Split(bearerToken, " ")
		if len(splitToken) == 2 {
			return splitToken[1]
		}
		return ""
	}

	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})
	return token, err
}
