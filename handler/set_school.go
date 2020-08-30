package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/class_schdule"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/class_schedule"
	"my_project/util"
)

type SetSchoolHandler struct {
	comm *model.CtxComm
	req  *class_schedule.SetSchoolsRequest
	resp *class_schedule.SetSchoolsResponse
}

func NewSetSchoolHandler(c *gin.Context) *SetSchoolHandler {
	resp := class_schedule.SetSchoolsResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	comm := util.GetCtxComm(c)
	req := genSetSchoolHandlerReq(c)
	return &SetSchoolHandler{
		req:  req,
		comm: comm,
		resp: &resp,
	}
}

func (h *SetSchoolHandler) Execute() *class_schedule.SetSchoolsResponse {
	if err := h.Check(); err != nil {
		logs.Error("SetSchoolHandler Execute err = %v", err)
		return h.resp
	}
	courseLoader := class_schdule.NewClassSchduler(h.comm.OpenId, h.req.SchoolName, h.req.StuId, h.req.StuPassword, h.req.VerifyCode)
	if courseLoader == nil {
		h.resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		h.resp.Message = "param illegal"
		h.resp.AlertMessage = "暂不支持该学校"
		return h.resp
	}
	err := courseLoader.Load()
	if err != nil {
		logs.Error("load course failed!err = %v", err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		h.resp.AlertMessage = courseLoader.GetErrorInfo()
		return h.resp
	}
	courses, err := courseLoader.GetCourse()
	if err != nil {
		logs.Error("GetCourse failed err = %v", err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		h.resp.AlertMessage = courseLoader.GetErrorInfo()
		return h.resp
	}
	err = h.saveCourse(courses)
	if err != nil {
		logs.Error("GetCourse failed err = %v", err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	err = db.SetSchool(h.comm.OpenId, h.req.SchoolName)
	if err != nil {
		logs.Error("set school err open_id = %v,err=%v", h.comm.OpenId, err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	err = db.SetStuSchool(h.comm.OpenId, h.req.SchoolName, h.req.StuId, h.req.StuPassword, courseLoader.GetFirsetWeekData())
	if err != nil {
		logs.Error("SetStuSchool err = %v", err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	return h.resp
}

func (h *SetSchoolHandler) Check() error {
	if h.comm == nil {
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param less"
		return errors.New("comm is nil")
	}
	if h.req == nil {
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param less"
		return errors.New("req is nil")
	}
	err := CheckSchoolIsLegal(h.req.SchoolName)
	if err != nil {
		h.resp.StatusCode = error_code.ERR_PARAM_NOT_FIND
		h.resp.Message = "param illegal"
		return fmt.Errorf("school  = %v is illegal", h.req.SchoolName)
	}
	return nil
}

func genSetSchoolHandlerReq(c *gin.Context) *class_schedule.SetSchoolsRequest {
	school := c.PostForm("school")
	stuId := c.PostForm("stu_id")
	stuPassword := c.PostForm("stu_password")
	code := c.PostForm("code")
	logs.Info("school = %v,stu_id = %v,stu_password = %v,code = %v", school, stuId, stuPassword, code)
	if school == "" {
		return nil
	}
	if stuId == "" {
		return nil
	}
	if stuPassword == "" {
		return nil
	}
	if code == "" {
		return nil
	}
	return &class_schedule.SetSchoolsRequest{
		SchoolName:  school,
		StuId:       stuId,
		StuPassword: stuPassword,
		VerifyCode:  code,
	}
}

func (h *SetSchoolHandler) saveCourse(courses [][][]*model.ClassInfo) error {
	for i, course := range courses {
		err := util.Retry(func() error {
			return db.SetCourse(h.comm.OpenId, h.req.SchoolName, h.req.StuId, i, course)
		}, 3)
		if err != nil {
			return err
		}
	}
	return nil
}
