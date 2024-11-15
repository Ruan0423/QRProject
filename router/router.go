package router

import (
	"web_framework/logger"
	"web_framework/settings"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"msg":settings.Conf.APP.Port,
		})
	})
	v1 := r.Group("api/v1")

	{
		v1.GET("/")
	}
	return r
}