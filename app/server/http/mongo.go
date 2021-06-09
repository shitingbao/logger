package http

import (
	"logger/app/model"
	"logger/lib/snsq"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setLog(ctx *gin.Context) {
	arg := new(snsq.NsqMes)
	if err := ctx.Bind(arg); err != nil {
		return
	}

	arg.Msg.Host = ctx.Request.RemoteAddr
	Sev.LogSeverNsq.LogPulish(arg)
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func LogFind(ctx *gin.Context) {
	arg := new(model.ArgLogCondition)
	if err := ctx.Bind(arg); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
		})
		return
	}
	res, err := Sev.Find(arg)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"data": res,
	})
}
