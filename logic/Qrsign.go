package logic

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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

func VerifyToken(token string) error {

	//验证token

	return redis.VerifyQrToken(token)
}

func generate_token() (string, error) {

	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// generateQr_url 生成qr
func generateQr_url(token string) (string, error) {

	//生成内容
	Qr_content := "http://127.0.0.1:8080/api/v1/qr/" + token

	fmt.Println("签到：",Qr_content)
	//二维码保存位置
	qr_root := "static/qrcodes/"
	qr_file_path := qr_root + token + ".png"

	defer delfile(qr_root,qr_file_path)
    // 创建二维码并保存
	err := qrcode.WriteFile(Qr_content, qrcode.Medium, 256, qr_file_path)
	if err != nil {
		zap.L().Error("创建二维码失败！", zap.Error(err))
		return "", err
	}

	return qr_file_path, nil

}

// delfile 删除多余的二维码
func delfile(dir string,newqr string)  {
	err:= filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if  err!=nil {
			return err
		}
		if !info.IsDir(){
			if filepath.Clean(path) != filepath.Clean(newqr) {
				zap.L().Info("我要删除二维码了！",zap.Any("遍历的path:",path),zap.Any("新生成的tokenqr:",newqr))
				return os.Remove(path)
			}
		}
		return nil
	})

	if err!=nil {
		zap.L().Error("删除二维码失败！",zap.Error(err))
		return
	}
}