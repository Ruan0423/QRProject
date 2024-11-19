package router

import (
	"web_framework/controller"
	"web_framework/logger"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {

		ctx.HTML(200, "index.html", "二维码签到")
	})
	v1 := r.Group("api/v1")

	{

		v1.GET("/getsignqr", controller.QrSignHandler)
	}
	return r
}
