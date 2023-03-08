package controllers

import (
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

type FindByIdReq struct {
	Id string `uri:"id" binding:"required"`
}

func (ec *EventController) FindEventById(ctx *gin.Context) {
	var req FindByIdReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ec.EventService.FindEventByName(req.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "FindEventById Failed", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type FindByNameReq struct {
	Name string `uri:"name" binding:"required"`
}

func (ec *EventController) FindEventByName(ctx *gin.Context) {
	var req FindByNameReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := ec.EventService.FindEventByName(req.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "FindEventByName Failed", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventController) AllEvents(ctx *gin.Context) {
	// FindAllAvaliableEvents() (*[]models.EventModel, error)

	result, err := ec.EventService.FindAllAvaliableEvents()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Find ALl Failed", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (ec *EventController) NewEvent(ctx *gin.Context) {
	var req models.EventModel

	bodyCheckError := middleware.CheckBodyBinding(&req, ctx)

	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError, "status": -1})
		return
	}

	result, err := ec.EventService.AddEvent(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "DB 저장 실패..", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, (result))
}

func (ec *EventController) RegisterEventRoutes(group *gin.RouterGroup) {
	group.GET("/findById/:id", ec.FindEventById)
	group.GET("/findByName/:name", ec.FindEventByName)
	group.GET("/", ec.AllEvents)
	group.POST("/", ec.NewEvent)
}
