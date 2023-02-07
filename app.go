package main

import (
	"tutorial/configs"
	"tutorial/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()

	v1 := router.Group("/v1")

	routes.MeterRoute(v1)
	routes.RecordRoute(v1)

	router.Run(":8080")
}
