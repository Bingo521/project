package db


import (
	"encoding/json"
	"my_project/model"
)

func GetCourseByWeekNum(openId string,schoolName string,weekNum int32)(*model.ClassRecord,error){
	classRecord := model.ClassRecord{}
	ndb := db.Where("open_id = ? and school_name = ? and week = ?",openId,schoolName,weekNum).First(&classRecord)
	if err := ndb.Error;err != nil{
		return nil, err
	}
	var class [][]*model.ClassInfo
	err:=json.Unmarshal([]byte(classRecord.Class),&class)
	if err != nil{
		return nil, err
	}
	classRecord.ClassInfos = class
	return &classRecord,nil
}

func SetCourse(openId string,schoolName string,stuId string,week int,courseInfo [][]*model.ClassInfo)error{
	classRecord := model.ClassRecord{}
	classRecord.OpenId = openId
	classRecord.SchoolName = schoolName
	classRecord.Week = int32(week)
	classInfo ,err:= json.Marshal(courseInfo)
	if err != nil{
		return err
	}
	classRecord.Class = string(classInfo)
	return db.Save(&classRecord).Error
}
