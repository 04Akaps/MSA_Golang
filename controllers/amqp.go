package controllers

import (
	"GO_MSA/services"

	"github.com/gin-gonic/gin"
)

type AmqpController struct {
	AmqpService services.Amqp
}

func NewAmqpController(amqp services.Amqp) AmqpController {
	return AmqpController{
		AmqpService: amqp,
	}
}

func (ac *AmqpController) TestGet(ctx *gin.Context) {
	ac.AmqpService.GetTest()
}

func (ac *AmqpController) RegisterAmqpRouter(server *gin.Engine) {
	amqpPath := server.Group("/amqp")

	amqpPath.GET("/", ac.TestGet)
}
