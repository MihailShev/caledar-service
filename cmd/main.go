package main

import (
	"github.com/MihailShev/caledar-service/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()
	router.POST("/add", controller.AddEvent)
	router.PUT("/update")

	err := router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
