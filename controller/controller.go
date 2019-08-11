package controller

import (
	"github.com/MihailShev/caledar-service/calendar"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
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
	val, err := strconv.ParseUint(uuid, 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid uuid %s", uuid)
	} else {
		event, ok := service.GetEventByUUID(val)

		if ok {
			response, err := proto.Marshal(&event)
			if err != nil {
				c.String(http.StatusInternalServerError, "%s", err.Error())
			} else {
				c.String(http.StatusOK, "%s", response)
			}
		} else {
			c.String(http.StatusNotFound, "Event with uuid: %s not found", uuid)
		}
	}
}

func UpdateHandler(c *gin.Context) {
	uuid, err := parseUUID(c)

	defer func() {
		if r := recover(); r != nil {
			c.String(http.StatusInternalServerError, "%s", r)
		}
	}()

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

	c.String(http.StatusOK, "")
}

func parseUUID(c *gin.Context) (uuid uint64, err error) {
	u := c.Param("uuid")
	val, err := strconv.ParseUint(u, 10, 64)

	return val, err
}

func parseEvent(c *gin.Context) (e *calendar.Event, err error) {
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(data, e)
	return
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
