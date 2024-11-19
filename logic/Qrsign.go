package logic

import (
	"crypto/rand"
	"encoding/base64"
	"web_framework/dao/redis"
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
