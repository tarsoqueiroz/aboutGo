package main

import (
	"restapi/db"
	"restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8910") // localhost:8910
}
