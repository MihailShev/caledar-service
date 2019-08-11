package main

import (
	"github.com/MihailShev/caledar-service/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()
	router.GET("/:uuid", controller.GetHandler)
	router.POST("/add", controller.AddHandler)
	router.PUT("/update", controller.UpdateHandler)

	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
