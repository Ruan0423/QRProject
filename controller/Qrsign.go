package controller

import (
	"encoding/base64"
	"web_framework/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// QrsignBybyteHandler 通过二维码的字节内容返回

func QrsignBybyteHandler(c *gin.Context) {
	data_byte, err := logic.Generate_Qr_By_byte()

	if err != nil {
		zap.L().Error("Generate_Qr_By_byte err:", zap.Error(err))
		return
	}

	base64data := base64.RawStdEncoding.EncodeToString(data_byte)

	c.JSON(200, gin.H{"data": base64data})
}

// QrSignHandler 生成图片url返回
func QrSignHandler(c *gin.Context) {

	qr_url, err := logic.Generate_Qr()
	if err != nil {
		c.JSON(10005, err)
		return
	}

	c.JSON(200, gin.H{
		"qrCodeUrl": "/" + qr_url,
	})
}

// DosignHandler 验证扫码的签到
func DosignHandler(c *gin.Context) {
	// 获取参数
	token := c.Param("token")
	// 验证token是否存在
	if err := logic.VerifyToken(token); err != nil {
		c.JSON(502, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{"msg": "签到成功"})
	}

}
