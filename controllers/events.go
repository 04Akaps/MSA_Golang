package controllers

import (
	"fmt"

	"GO_MSA/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService services.Event
}

func NewEventController(event services.Event) EventController {
	return EventController{
		EventService: event,
	}
}

func (ec *EventController) FindEvent(ctx *gin.Context) {
	fmt.Println("Find Event")
}

func (ec *EventController) AllEvents(ctx *gin.Context) {
	fmt.Println("Find ALl Event")
}

func (ec *EventController) newEvent(ctx *gin.Context) {
	fmt.Println("Find ALl New")
}

func (ec *EventController) RegisterEventRoutes(group *gin.RouterGroup) {
	group.GET("/:SearchCriterial/:search", ec.FindEvent)
	group.GET("/", ec.AllEvents)
	group.POST("/", ec.newEvent)
}
