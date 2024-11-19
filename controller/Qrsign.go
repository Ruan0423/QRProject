package controller

import (
	"web_framework/logic"

	"github.com/gin-gonic/gin"
)

func QrSignHandler(c *gin.Context) {

	qr_url, err := logic.Generate_Qr()
	if err != nil {
		c.JSON(10005, err)
		return
	}
	c.JSON(200, gin.H{
		"qrCodeUrl": qr_url,
	})
}
