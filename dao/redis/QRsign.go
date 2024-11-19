package redis

import (
	"context"
	"time"
)

func QRsign(token string) error {
	err:=Rdb.Set(context.Background(), "Sign_Token", token, 10*time.Second).Err()
	if err!=nil {
		return err
	}
	return nil
}