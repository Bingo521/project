package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/class_schedule"
	"my_project/util"
)

type SetSchoolHandler struct {
	comm *model.CtxComm
	req *class_schedule.SetSchoolsRequest
	resp *class_schedule.SetSchoolsResponse
}

func NewSetSchoolHandler(c *gin.Context)*SetSchoolHandler{
	resp := class_schedule.SetSchoolsResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	comm := util.GetCtxComm(c)
	req := genSetSchoolHandlerReq(c)
	return &SetSchoolHandler{
		req: req,
		comm: comm,
		resp: &resp,
	}
}

func (h *SetSchoolHandler)Execute()*class_schedule.SetSchoolsResponse{
	if err := h.Check();err != nil{
		logs.Error("SetSchoolHandler Execute err = %v",err)
		return h.resp
	}
	err:=db.SetSchool(h.comm.OpenId,h.req.SchoolName)
	if err != nil{
		logs.Error("set school err open_id = %v,err=%v",h.comm.OpenId,err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		return h.resp
	}
	return h.resp
}

func (h *SetSchoolHandler)Check()error{
	if h.comm == nil{
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param less"
		return errors.New("comm is nil")
	}
	if h.req == nil{
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param less"
		return errors.New("req is nil")
	}
	err:=CheckSchoolIsLegal(h.req.SchoolName)
	if err != nil{
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param illegal"
		return fmt.Errorf("school  = %v is illegal",h.req.SchoolName)
	}
	return nil
}

func genSetSchoolHandlerReq(c *gin.Context)(*class_schedule.SetSchoolsRequest){
	school:=c.PostForm("school")
	if school == ""{
		return nil
	}
	return &class_schedule.SetSchoolsRequest{
		SchoolName: school,
	}
}