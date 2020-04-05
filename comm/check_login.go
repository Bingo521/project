package comm

import (
	"github.com/gin-gonic/gin"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/util"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if path == "/login" {
			c.Next()
			return
		}
		sessionKey, err := c.Request.Cookie("session")
		if err != nil {
			resp := util.MakeErrResp(error_code.ERR_NOT_LOGIN, "please login")
			c.JSON(200, resp)
			c.Abort()
			return
		}
		openId, err := c.Request.Cookie("openid")
		if err != nil {
			resp := util.MakeErrResp(error_code.ERR_NOT_LOGIN, "please login")
			c.JSON(200, resp)
			c.Abort()
			return
		}
		logs.Info("session = %v,openid = %v", sessionKey, openId)
		loginInfo, err := redis.GetSession(sessionKey.Value)
		if err != nil {
			logs.Error("get session err key = %v,err = %v", sessionKey, err)
			resp := util.MakeErrResp(error_code.ERR_NOT_LOGIN, "please login")
			c.JSON(200, resp)
			c.Abort()
			return
		}
		if openId.Value != loginInfo.OpenId {
			logs.Error("openId = %v != session openid = %v", openId, loginInfo.OpenId)
			resp := util.MakeErrResp(error_code.ERR_NOT_LOGIN, "please login")
			c.JSON(200, resp)
			c.Abort()
			return
		}
		redis.FlushSession(loginInfo.SessionKey)
		logs.Info("check success!")
		c.Next()
	}
}
