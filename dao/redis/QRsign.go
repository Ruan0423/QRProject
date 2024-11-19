package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

func QRsign(token string) error {
	err := Rdb.Set(context.Background(), "Sign_Token", token, 10*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func VerifyQrToken(token string) error {
	//验证token是否存在和是否过期
	result, err := Rdb.Get(context.Background(), "Sign_Token").Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("二维码已过期")
		}
		return err
	}

	if result != token {
		return errors.New("二维码不存在")
	}
	return nil

}
