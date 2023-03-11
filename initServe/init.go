package initServe

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ar httpsServer *http.Server
// ar server *gin.Engine
func ServeAPI(httpPath string, httpsServer *http.Server, httpServer *gin.Engine) (chan error, chan error) {
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)

	go func() {
		httpErrChan <- http.ListenAndServe(":80", httpServer)
	}()

	go func() {
		httpsErrChan <- httpsServer.ListenAndServeTLS("", "")
	}()

	return httpErrChan, httpsErrChan
}

func SetHttpsServer(httpsServer *gin.Engine) {
	trustAddress := []string{"http://127.0.0.1"} // nginx를 사용할 예정이기 떄문에
	httpsServer.SetTrustedProxies(trustAddress)
	httpsServer.Use(gin.Logger())   // Logger를 통해서 DefaultWriter에 로그를 기록 -> 로그 형태 변환
	httpsServer.Use(gin.Recovery()) // panic이 발생하면 500에러

	// Set Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = trustAddress // 모든 도메인에 대한 요청을 허용
	httpsServer.Use(cors.New(config))
}

func SetHttpServer(httpsServer *gin.Engine) {
}
