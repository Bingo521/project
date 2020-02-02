package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"my_project/conf"
)

var RedisClient redis.Conn
// 初始化redis链接池
func init() {
	redisConf:=conf.Conf.GetRedis()
	var err error
	RedisClient, err = redis.Dial(redisConf.Network, redisConf.Ip+":"+redisConf.Port)
	if err != nil {
		panic(fmt.Sprintf("redis connect err:%v",err))
	}
}
