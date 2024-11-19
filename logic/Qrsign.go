package logic

import (
	"crypto/rand"
	"encoding/base64"
	"web_framework/dao/redis"

	"github.com/skip2/go-qrcode"
	"go.uber.org/zap"
)

func Generate_Qr() (sign_url string, err error) {
	//1. 生成二维码token
	token, err := generate_token()
	if err != nil {
		return "", err
	}
	//2.使用redis存储token

	err = redis.QRsign(token)

	//3. 调用第三方 API生成二维码并返回Url
	sign_url, err = generateQr_url(token)

	return
}

func generate_token() (string, error) {

	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func generateQr_url(token string) (string, error) {

	//生成内容
	Qr_content := "http://127.0.0.1:8080/qr/" + token
	//二维码保存位置
	// root_path := "static/qrcodes"
	qr_file_path := "static/qrcodes/" + token + ".png"

	err := qrcode.WriteFile(Qr_content, qrcode.Medium, 256, qr_file_path)
	if err != nil {
		zap.L().Error("创建二维码失败！", zap.Error(err))
		return "", err
	}

	return "/" + qr_file_path, nil

}
