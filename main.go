package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"GO_MSA/config"
	"GO_MSA/controllers"
	"GO_MSA/initServe"
	"GO_MSA/mongo"
	"GO_MSA/services"

	"github.com/gin-gonic/gin"
)

var (
	httpsServer *gin.Engine
	httpServer  *gin.Engine
)

var (
	// event variable
	eventCtx        context.Context
	eventService    services.Event
	eventController controllers.EventController
)

var ctxMongo context.Context

var envConfig config.Config

func init() {
	envConfig = config.LoadConfig(".")
	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	gin.EnableJsonDecoderDisallowUnknownFields()

	httpsServer = gin.New()
	httpServer = gin.New()

	initServe.SetHttpsServer(httpsServer) // httpsServer 기본 셋티
	initServe.SetHttpServer(httpServer)   // httpServer 기본 셋팅
	// Set Event Controller
	eventCtx = context.Background()
}

func main() {
	// mongo Session
	mongoDBLayout, err := mongo.NewMongoSession(envConfig)
	if err != nil {
		log.Fatal("Mongo Session Connection Error", err)
	}

	eventService = services.NewEventService(eventCtx, mongoDBLayout)
	eventController = controllers.NewEventController(eventService)
	eventController.RegisterEventRoutes(httpsServer)

	tlsConfig, err := config.GetTlsConfig(envConfig)
	if err != nil {
		log.Fatal("GetTlsConfig Error", err)
	}

	// HTTPS 설정
	httpsServer := &http.Server{
		Addr:      envConfig.ServerAddress,
		Handler:   httpsServer,
		TLSConfig: tlsConfig,
	}

	// MSA구조 이기 떄문에 또 하나의 Server를 열어보자
	// 기존 HTTPS서버는 gin을 사용하고 있기 떄문에 이번에는 net/http로 그냥 http서버만 open할 예정
	// 하지만 ListenAndServeTLS는 시작하면 둘다 함수를 중단시킨다.
	// 그러기 떄문에 다른 Go Routin으로 실행시키고 해당 에러에 대해서 channels를 만들어서 구성하자

	httpErrChan, httpsErrChan := initServe.ServeAPI(":80", httpsServer, httpServer)

	// http와 https로 들어오는 서버 중 에러를 발생시키는 case를 탐지
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Server Error", err)
	case err := <-httpsErrChan:
		log.Fatal("HTTPS Server Error", err)
	}

	fmt.Println("Server is Started")

	defer mongoDBLayout.Session.Disconnect(ctxMongo) // 리소스를 줄이기 위해서 mongo에 대한 Close를 defer로 호출
}
