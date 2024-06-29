package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heyrmi/go-webapi/db"
	"github.com/heyrmi/go-webapi/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
