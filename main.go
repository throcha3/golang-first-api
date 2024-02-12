package main

import (
	routes "api/api/routes"
	"api/database/mongodb"

	"github.com/gin-gonic/gin"
)

func main(){
	app := gin.Default()

	routes.AppRoutes(app)

	mongodb.InitConn()

	app.Run("localhost:3001")
}