package db

import (
	"my_project/model"
)

func GetStuSchool(openId string,schoolName string)(*model.StuSchool,error){
	stuSchool:=model.StuSchool{}
	ndb := db.Where("open_id = ? and school_name = ?",openId,schoolName).First(&stuSchool)
	if err:=ndb.Error;err != nil{
		return nil, err
	}
	return &stuSchool,nil
}

func SetStuSchool(openId string,schoolName string,stuId string,stuPassword string,date int64)error{
	stuInfo := new(model.StuSchool)
	stuInfo.OpenId = openId
	stuInfo.SchoolName = schoolName
	stuInfo.StuId = stuId
	stuInfo.StuPassword = stuPassword
	stuInfo.FirstWeekDate = date
	return db.Save(stuInfo).Error
}

func SetStuSchoolFirstWeekTime(openId string,schoolName string,stuId string,date int64)error{
	stuInfo := new(model.StuSchool)
	stuInfo.OpenId = openId
	stuInfo.SchoolName = schoolName
	stuInfo.StuId = stuId
	return db.Where("open_id = ? and school_name = ? and stu_id = ?",openId,schoolName,stuId).Update("first_week_date",date).Error
}
