package controller

import (
	"errors"
	"fmt"
	"github.com/MihailShev/caledar-service/calendar"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
)

var service = calendar.NewCalendar()

func AddHandler(c *gin.Context) {
	event, err := parseEvent(c)

	if err != nil {
		c.String(http.StatusBadRequest, "%s", err)
		return
	}

	addedEvent := service.AddEvent(*event)

	response, err := proto.Marshal(&addedEvent)

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	c.String(http.StatusCreated, "%s", response)
}

func GetHandler(c *gin.Context) {
	uuid, err := parseUUID(c)

	if err != nil {
		c.String(http.StatusBadRequest, "%s", err)
		return
	}

	event, ok := service.GetEventByUUID(uuid)

	if !ok {
		c.String(http.StatusNotFound, "Event with uuid: %s not found", uuid)
		return
	}

	response, err := proto.Marshal(&event)

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	c.String(http.StatusOK, "%s", response)
}

func UpdateHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.String(http.StatusInternalServerError, "%s", r)
		}
	}()

	uuid, err := parseUUID(c)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid uuid %s", uuid)
		return
	}

	src, err := parseEvent(c)

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
		return
	}

	event, ok := service.GetEventByUUID(uuid)

	if !ok {
		c.String(http.StatusNotFound, "Event with uuid: %s not found", uuid)
		return
	}

	proto.Merge(&event, src)

	err = service.ReplaceEvent(event)

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
		return
	}

	response, err := proto.Marshal(&event)

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
	}

	c.String(http.StatusOK, "%s", response)
}

func parseUUID(c *gin.Context) (uuid uint64, err error) {
	u := c.Param("uuid")
	val, err := strconv.ParseUint(u, 10, 64)

	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid uuid %s, err:%s", u, err.Error()))
	}

	return val, err
}

func parseEvent(c *gin.Context) (event *calendar.Event, err error) {
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return nil, err
	}

	event = &calendar.Event{}
	err = proto.Unmarshal(data, event)

	return
}
