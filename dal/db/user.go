package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"my_project/logs"
	"my_project/model"
)

func GetUserInfo(openId string)(*model.UserInfo,error){
	user := model.UserInfo{}
	ndb:=db.Where("open_id = ?", openId).First(&user)
	if ndb.Error != nil{
		return nil,ndb.Error
	}
	return &user,nil
}

func SetSchool(openId string,schoolName string)error{
	userInfo,err:=GetUserInfo(openId)
	logs.Error("get user open_id = %v,err=%v",openId,err)
	if err != nil{
		if err.Error() == gorm.ErrRecordNotFound.Error(){
			userInfo = new(model.UserInfo)
			userInfo.OpenId = openId
			userInfo.CurSchoolName =schoolName
			logs.Info("save user = %v",*userInfo)
			return db.Create(userInfo).Error
		}
		return err
	}
	if userInfo == nil{
		return fmt.Errorf("find user info is nil")
	}
	logs.Info("user info = %v",*userInfo)
	ndb := db.Model(&userInfo).Update("cur_school_name",schoolName)
	return ndb.Error
}

