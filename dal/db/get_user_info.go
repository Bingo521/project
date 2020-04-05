package db

import "my_project/model"

func GetUserInfo(openId string)(*model.UserInfo,error){
	user := model.UserInfo{}
	db=db.Where("open_id = ?", openId).First(&user)
	if db.Error != nil{
		return nil,db.Error
	}
	return &user,nil
}
