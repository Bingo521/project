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
	"strconv"
)

type GetCourseByWeekNumHandler struct {
	c *gin.Context
	comm *model.CtxComm
	req *class_schedule.GetCourseByWeekNumRequest
	resp *class_schedule.GetCourseByWeekNumResponse
}

func NewGetCourseByWeekNumHandler(c *gin.Context)*GetCourseByWeekNumHandler{
	resp := class_schedule.GetCourseByWeekNumResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	return &GetCourseByWeekNumHandler{
		c: c,
	}
}

func (h *GetCourseByWeekNumHandler)Execute()*class_schedule.GetCourseByWeekNumResponse{
	if conf.Conf.IsDebug(){
		return makeGetCourseByWeekDebugResp()
	}
	var err error
	err = h.makeReq()
	if err != nil{
		logs.Error("make req err:%v",err)
		h.resp.AlertMessage = "请求异常"
		return h.resp
	}
	class,err:=h.getClass()
	if err != nil{
		logs.Error("get class failed! open_id = %v,err = %v",h.comm.OpenId,err)
		h.resp.StatusCode = error_code.ERR_SERVER_ERR
		h.resp.Message = "service err"
		h.resp.AlertMessage = "服务异常"
		return h.resp
	}
	h.resp.SchoolName = h.req.SchoolName
	h.resp.Courses = class
	return h.resp
}

func (h* GetCourseByWeekNumHandler)makeReq()error{
	h.comm = util.GetCtxComm(h.c)
	if h.comm  == nil{
		h.resp.StatusCode = error_code.ERR_NOT_LOGIN
		h.resp.Message = "please login"
		return fmt.Errorf("make comm fail")
	}
	strWeekNum := h.c.Query("week_num")
	weekNum,err:=strconv.ParseInt(strWeekNum,10,64)
	if err != nil{
		h.resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		h.resp.Message = "param illegal"
		return fmt.Errorf("week_num = [%v] is illegal",strWeekNum)
	}
	schoolName := h.c.Query("school")
	if schoolName == ""{
		h.resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		h.resp.Message = "param illegal"
		return fmt.Errorf("schoolName = [%v] is illegal",schoolName)
	}
	h.req = &class_schedule.GetCourseByWeekNumRequest{
		Week: int32(weekNum),
		SchoolName: schoolName,
	}
	return nil
}

func (h* GetCourseByWeekNumHandler)getClass()([]*class_schedule.Course,error){
	record,err:=db.GetCourseByWeekNum(h.comm.OpenId,h.req.SchoolName,h.req.Week)
	if err != nil{
		return nil, err
	}
	return GetClass(record.ClassInfos),nil
}

func makeGetCourseByWeekDebugResp()*class_schedule.GetCourseByWeekNumResponse{
	resp := &class_schedule.GetCourseByWeekNumResponse{}
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