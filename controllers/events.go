package controllers

import (
	"fmt"
	"net/http"

	"GO_MSA/middleware"
	"GO_MSA/models"
	"GO_MSA/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService services.Event
}

type ReqTest struct {
	NameTwo   string `json:"name_two" binding:"required"`
	NameThree string `json:"name_three" binding:"required"`
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

func (ec *EventController) NewEvent(ctx *gin.Context) {
	var req models.EventModel

	bodyCheckError := middleware.CheckBodyBinding(&req, ctx)

	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError, "status": -1})
		return
	}

	fmt.Println("Find ALl New")
}

func (ec *EventController) RegisterEventRoutes(group *gin.RouterGroup) {
	group.GET("/:SearchCriterial/:search", ec.FindEvent)
	group.GET("/", ec.AllEvents)
	group.POST("/", ec.NewEvent)
}
