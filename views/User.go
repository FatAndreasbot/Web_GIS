package views

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func User_endpoints(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
