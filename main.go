package main

import (
	"config"
	"views"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.DB_init()

	views.Player_CRUD_endpoints(r)
}
