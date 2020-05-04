package class_schdule

import "my_project/model"

type ClassSchduler interface {
	Load()error
	GetCourse()([][][]*model.ClassInfo,error)
	GetFirsetWeekData()int64
	GetErrorInfo()string
}

type NewSchduler func(openId string,schoolName string,stuId string,stuPassword string,code string)ClassSchduler
var schduleMap = map[string]NewSchduler{
	"亳州学院":GetBoZhouLoader,
}

func NewClassSchduler(openId string,schoolName string,stuId string,stuPassword string,code string) ClassSchduler {
	f,find:=schduleMap[schoolName]
	if find && f != nil{
		return f(openId,schoolName,stuId,stuPassword,code)
	}
	return nil
}
