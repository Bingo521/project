package redis

import (
	"github.com/garyburd/redigo/redis"
	"my_project/logs"
	"time"
)

func Get(key string) (string, error) {
	resp, err := RedisClient.Do("get", key)
	if err != nil {
		return "", err
	}
	logs.Info("Get %v", resp)
	str, ok := resp.([]byte)
	if !ok {
		logs.Error("resp = %v is not string type!", resp)
		return "", err
	}
	return string(str), nil
}

func Set(key string, val string) error {
	_, err := RedisClient.Do("set", key, val)
	return err
}

func SetTimeOut(key string, second time.Duration) error {
	_, err := RedisClient.Do("expire", key, second)
	return err
}

func Inc(key string) (int64, error) {
	RedisClient.Do("")
	v, err := redis.Int64(RedisClient.Do("INCR", key))
	if err != nil {
		return 0, err
	}
	return v, nil
}
