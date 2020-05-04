package handler

import (
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/class_schedule"
	"my_project/util"
	"strconv"
)

type SetWeekHandler struct {
	c *gin.Context
	comm *model.CtxComm
}

func NewSetWeekHandler(c *gin.Context)*SetWeekHandler{
	comm := util.GetCtxComm(c)
	return &SetWeekHandler{
		c: c,
		comm: comm,
	}
}

func (h*SetWeekHandler)Execute()*class_schedule.SetCurWeekResponse{
	resp := class_schedule.SetCurWeekResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	if h.c == nil || h.comm == nil{
		resp.StatusCode = error_code.ERR_NOT_LOGIN
		resp.Message = "please login"
		return &resp
	}

	week := h.c.PostForm("week")
	iWeek,err:=strconv.ParseInt(week,10,64)
	if err != nil{
		logs.Error("week = %v is illegal!",week)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = "param illegal"
		return &resp
	}
	userInfo,err:=db.GetUserInfo(h.comm.OpenId)
	if err != nil{
		logs.Error("user info load failed ")
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = "service err"
		return &resp
	}
	_=iWeek
	_=userInfo
	return &resp
}
