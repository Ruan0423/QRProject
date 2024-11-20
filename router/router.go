package router

import (
	"web_framework/controller"
	"web_framework/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 启用 CORS 支持
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080", "http://localhost:8080"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST"},                                    // 允许的方法
		AllowHeaders:     []string{"Content-Type"},                                   // 允许的头
		AllowCredentials: true,                                                       // 是否允许跨域请求携带凭证
	}))
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	{
		r.GET("/", func(ctx *gin.Context) {

			ctx.HTML(200, "index.html", "二维码签到")
		})

		r.GET("/qrbyte", func(ctx *gin.Context) {

			ctx.HTML(200, "qrbyte.html", "二维码签到")
		})
	}

	v1 := r.Group("api/v1")

	{
		v1.GET("/getsignqr", controller.QrSignHandler)
		v1.GET("/qr/:token", controller.DosignHandler)
		v1.GET("/getQRCodebyte", controller.QrsignBybyteHandler)
	}
	return r
}
