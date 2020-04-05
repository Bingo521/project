package comm

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"my_project/conf"
)

func SetSession()gin.HandlerFunc{
	redisConf:= conf.Conf.GetRedis()
	store, _ := redis.NewStore(10,redisConf.Network , redisConf.Ip+":"+redisConf.Port, redisConf.Password, []byte("gdhasdq82bhsadb,.aswqioh"))
	return  sessions.Sessions("mysession", store)
}
