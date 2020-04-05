package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_project/dal/cache"
	"my_project/error_code"
	"my_project/logs"
	"my_project/proto_gen/class_schedule"
)

type GetSchoolHandler struct {
}

func NewGetSchoolHandler() *GetSchoolHandler {
	return &GetSchoolHandler{}
}

func (h *GetSchoolHandler) Execute(ctx *gin.Context) *class_schedule.GetSchoolsResponse {
	resp := class_schedule.GetSchoolsResponse{}
	resp.StatusCode = error_code.ERR_SUCCESS
	resp.Message = "success"
	resp.Schools = cache.GetSchools()
	b, _ := json.Marshal(resp)
	logs.Info("get school = resp=%v", string(b))
	return &resp
}
