package db

import "my_project/model"

func GetStuSchool(openId string,schoolName string)(*model.StuSchool,error){
	stuSchool:=model.StuSchool{}
	db = db.Where("open_id = ? and school_name = ?",openId,schoolName).First(&stuSchool)
	if err:=db.Error;err != nil{
		return nil, err
	}
	return &stuSchool,nil
}
