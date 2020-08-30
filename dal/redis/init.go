package redis

import (
	"github.com/garyburd/redigo/redis"
	"my_project/conf"
	"my_project/logs"
	"my_project/model"
	"strings"
	"sync"
	"time"
)

var RedisClient *RedisConn

// 初始化redis链接池
func init() {
	redisConf := conf.Conf.GetRedis()
	RedisClient = NewRedisConn(redisConf)
	if err := RedisClient.Conn(); err != nil {
		panic(err)
	}
	RedisClient.KeepLive()
}

type RedisConn struct {
	redisClient redis.Conn
	redisConf   *model.Redis
	mu          sync.Mutex
}

func NewRedisConn(redisConf *model.Redis) *RedisConn {
	redisConn := &RedisConn{
		redisConf: redisConf,
		mu:        sync.Mutex{},
	}
	return redisConn
}

func (r *RedisConn) Conn() error {
	r.mu.Lock()
	if r.redisClient != nil {
		if _, err := r.redisClient.Do("ping"); err == nil {
			return nil
		}
	}
	logs.Info("redis ip = %v,port = %v", r.redisConf.Ip, r.redisConf.Port)
	redisClient, err := redis.Dial(r.redisConf.Network, r.redisConf.Ip+":"+r.redisConf.Port)
	if err != nil {
		return err
	}
	err = redisClient.Send("auth", r.redisConf.Password)
	if err != nil {
		return err
	}
	r.redisClient = redisClient
	r.mu.Unlock()
	return nil
}

func (r *RedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	reply, err = r.redisClient.Do(commandName, args...)
	if err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			r.tryConnWithRetry(3)
		}
		return reply, err
	}
	return reply, err
}

func (r *RedisConn) KeepLive() {
	go func() {
		var err error
		for {
			if err = r.Ping(); err != nil {
				r.tryConnWithRetry(3)
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (r *RedisConn) Ping() error {
	var err error
	for i := 0; i < 3; i++ {
		if err = r.redisClient.Send("ping"); err == nil {
			return nil
		}
	}
	return err
}

func (r *RedisConn) tryConnWithRetry(limiRetryTime int) error {
	var err error
	for i := 0; i < limiRetryTime; i++ {
		if err = r.Conn(); err == nil {
			return nil
		}
	}
	return err
}
