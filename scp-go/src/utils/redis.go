package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetKeyRDB(rdb *redis.Client, key string) (string, error) {
	ctxGet, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	strCmd := rdb.Get(ctxGet, key)
	if strCmd.Err() != nil {
		return "", strCmd.Err()
	}
	res := strCmd.Val()
	return res, nil
}

func SetKeyRDB(rdb *redis.Client, key string, val string, timeout time.Duration) error {
	ctxSetRDB, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	statCmdSetIDAT := rdb.Set(ctxSetRDB, key, val, timeout)
	return statCmdSetIDAT.Err()
}
