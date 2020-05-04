package handler

import (
	"github.com/gin-gonic/gin"
	"my_project/error_code"
	"my_project/proto_gen/class_schedule"
	"my_project/verification_code"
)

type GetVerifyCode struct {
	c *gin.Context
}
func NewGetVerifyCodeHandler(c *gin.Context)* GetVerifyCode{
	return &GetVerifyCode{
		c: c,
	}
}

func (h *GetVerifyCode)Execute()*class_schedule.GetVerifyCodeResponse{
	resp := class_schedule.GetVerifyCodeResponse{}
	resp.Message = "success"
	resp.StatusCode = 0
	schoolName := h.c.Query("school_name")
	if schoolName == ""{
		resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		resp.Message = "param less"
		return &resp
	}
	link:=verification_code.GetVerifyCodeLinkBySchoolName(schoolName)
	resp.Url = link
	return &resp
}
