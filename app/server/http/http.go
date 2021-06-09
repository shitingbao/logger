package http

import (
	"fmt"
	"logger/app/service"

	"github.com/gin-gonic/gin"
)

var Sev *service.Service

func HttpInit(s *service.Service) {
	Sev = s
	r := gin.Default()
	engine(r)
	r.Run(fmt.Sprintf("%s:%s", s.Config.HTTPSvc.Host, s.Config.HTTPSvc.Port))
}

func engine(r *gin.Engine) {
	r.POST("/log/set", setLog)
	r.POST("/log/find", LogFind)
}
