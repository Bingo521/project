package handler

import (
	"github.com/gin-gonic/gin"
	"my_project/proto_gen/login"
)

type Login struct {
	ctx *gin.Context
}

func NewLogin(c *gin.Context)*Login{
	return &Login{
		ctx:c,
	}
}

func (h *Login)Execute()*login.LoginResponse{
	return &login.LoginResponse{
		StatusCode:0,
		Message:"success",
		Openid:"123456",
		SessionKey:"dsha",
	}
}