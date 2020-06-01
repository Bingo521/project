package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/conf"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/class_schedule"
	"my_project/util"
	"time"
)

type GetCurCurrentCourseHandler struct {
	c *gin.Context
	comm *model.CtxComm
	resp *class_schedule.GetCurClassResponse
}

func NewGetCurCurrentCourseHandler(c *gin.Context)*GetCurCurrentCourseHandler{
	resp := &class_schedule.GetCurClassResponse{
		StatusCode: 0,
		Message: "success",
	}
	return &GetCurCurrentCourseHandler{
		c: c,
		resp: resp,
	}
}

func (h *GetCurCurrentCourseHandler)Execute()*class_schedule.GetCurClassResponse{
	err := h.makeReq()
	if err != nil{
		logs.Error("make req err=%v",err)
		return h.resp
	}
	if conf.Conf.IsDebug(){
		return debugResp()
	}
	schoolName,err := h.getCurSchoolName()
	if err != nil{
		logs.Error("get school name err:%v",err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	weekNum ,err:= h.getWeekNum(schoolName)
	if err != nil{
		logs.Error("get week num err:%v",err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	record,err:=db.GetCourseByWeekNum(h.comm.OpenId,schoolName,weekNum)
	if err != nil || record == nil{
		logs.Error("GetCourseByWeekNum openId = %v,schoolName = %v,weekNum = %v,err = %v",h.comm.OpenId,schoolName,weekNum)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
	}
	class := GetClass(record.ClassInfos)
	h.resp.Courses = class
	return h.resp
}

func (h *GetCurCurrentCourseHandler)makeReq()error{
	h.comm = util.GetCtxComm(h.c)
	if h.comm == nil{
		h.resp.StatusCode = error_code.ERR_NOT_LOGIN
		h.resp.Message = "please login"
		h.resp.AlertMessage = "请先登录"
		return fmt.Errorf("make comm err open_id = %v",h.comm.OpenId)
	}
	return nil
}

func (h *GetCurCurrentCourseHandler)getCurSchoolName()(string,error){
	userInfo,err:=db.GetUserInfo(h.comm.OpenId)
	if err != nil || userInfo == nil{
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return "",err
	}
	return userInfo.CurSchoolName,nil
}

func (h* GetCurCurrentCourseHandler)getWeekNum(schoolName string)(int32,error){
	stuSchool,err:=db.GetStuSchool(h.comm.OpenId,schoolName)
	if err != nil{
		return 0, err
	}
	firstWeekDate:=time.Unix(stuSchool.FirstWeekDate,0)
	now := time.Now()
	subDay := now.Day() - firstWeekDate.Day()
	subWeek := subDay / 7
	if subDay % 7 != 0{
		subWeek ++
	}
	return int32(subWeek),nil
}

func debugResp()*class_schedule.GetCurClassResponse{
	resp := &class_schedule.GetCurClassResponse{}
	resp.Message = "success"
	resp.StatusCode = 0
	resp.SchoolName = "北京大学"
	resp.Week = 1
	resp.Courses = []*class_schedule.Course{
		&class_schedule.Course{
			WeekNum: class_schedule.Week_Mon,
			ClassNum:1,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},
		&class_schedule.Course{
			WeekNum: class_schedule.Week_Mon,
			ClassNum:2,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: class_schedule.Week_Tue,
			ClassNum:3,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: class_schedule.Week_Tue,
			ClassNum:4,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 3,
			ClassNum:1,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 3,
			ClassNum:4,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 4,
			ClassNum:5,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 4,
			ClassNum:6,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 4,
			ClassNum:3,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 5,
			ClassNum:1,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 6,
			ClassNum:1,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},&class_schedule.Course{
			WeekNum: 7,
			ClassNum:1,
			CourseName:"高等数学",
			TeacherName:"张翠芳",
			Place: "1号楼B432",
		},
	}
	return resp
}