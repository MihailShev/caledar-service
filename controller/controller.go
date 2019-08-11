package controller

import (
	"fmt"
	"github.com/MihailShev/caledar-service/calendar"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

var service = calendar.NewCalendar()

func AddHandler(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	var response []byte

	if err == nil {
		response, err = addEvent(data)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
	} else {
		c.String(http.StatusCreated, "%s", response)
	}
}

func GetHandler(c *gin.Context) {
	uuid := c.Param("uuid")
	fmt.Println("uuid", uuid)
}

func addEvent(message []byte) (response []byte, err error) {
	event := calendar.Event{}

	err = proto.Unmarshal(message, &event)

	if err != nil {
		return
	}

	event = service.AddEvent(event)
	fmt.Println("add event", event)
	response, err = proto.Marshal(&event)

	return
}
