package util

import (
	"github.com/gin-gonic/gin"
	"my_project/logs"
	"my_project/model"
)

func GetCtxComm(c *gin.Context)*model.CtxComm{
	sessionKey, err := c.Request.Cookie("session")
	if err != nil {
		logs.Error("not find session key")
		return nil
	}
	openId, err := c.Request.Cookie("openid")
	if err != nil {
		logs.Error("not find open id")
		return nil
	}
	return &model.CtxComm{
		SessionId: sessionKey.Value,
		OpenId: openId.Value,
	}
}
