package db

import (
	"encoding/json"
	"my_project/model"
)

func GetCourseByWeekNum(openId string,schoolName string,weekNum int32)(*model.ClassRecord,error){
	classRecord := model.ClassRecord{}
	db = db.Where("open_id = ? and school_name = ? and week = ?",openId,schoolName,weekNum).First(&classRecord)
	if err := db.Error;err != nil{
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
