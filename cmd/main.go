package main

import (
	"github.com/MihailShev/caledar-service/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()
	router.GET("/:uuid", controller.GetHandler)
	router.POST("/:uuid", controller.UpdateHandler)
	router.POST("/", controller.AddHandler)

	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
