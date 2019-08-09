package controller

import (
	"github.com/MihailShev/caledar-service/calendar"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

var service = &calendar.Calendar{}

func AddEvent(c *gin.Context) {
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

func addEvent(message []byte) (response []byte, err error) {
	event := calendar.Event{}

	err = proto.Unmarshal(message, &event)

	if err != nil {
		return
	}

	event = service.AddEvent(event)
	response, err = proto.Marshal(&event)

	return
}
