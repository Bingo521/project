package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"my_project/logs"
	"my_project/model"
)

func GetUserInfo(openId string) (*model.UserInfo, error) {
	user := model.UserInfo{}
	ndb := db.Model(model.UserInfo{}).LogMode(true).Where("open_id = ?", openId).First(&user)
	if ndb.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if ndb.Error != nil {
		return nil, ndb.Error
	}
	return &user, nil
}

func MGetUserInfo(openIds []string) (map[string]*model.UserInfo, error) {
	distinctOpenIds := make([]string, 0, len(openIds))
	openIdMap := make(map[string]bool, len(openIds))
	for _, openId := range openIds {
		openIdMap[openId] = true
	}
	for openId, _ := range openIdMap {
		distinctOpenIds = append(distinctOpenIds, openId)
	}
	users := []model.UserInfo{}
	ndb := db.Model(model.UserInfo{}).LogMode(true).Where("open_id in (?)", distinctOpenIds).Find(&users)
	if ndb.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if ndb.Error != nil {
		return nil, ndb.Error
	}
	openId2UserInfo := make(map[string]*model.UserInfo, len(users))
	for _, userInfo := range users {
		openId2UserInfo[userInfo.OpenId] = &userInfo
	}
	return openId2UserInfo, nil
}

func SetSchool(openId string, schoolName string) error {
	userInfo, err := GetUserInfo(openId)
	logs.Error("get user open_id = %v,err=%v", openId, err)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			userInfo = new(model.UserInfo)
			userInfo.OpenId = openId
			userInfo.CurSchoolName = &schoolName
			logs.Info("save user = %v", *userInfo)
			return db.Create(userInfo).Error
		}
		return err
	}
	if userInfo == nil {
		return fmt.Errorf("find user info is nil")
	}
	logs.Info("user info = %v", *userInfo)
	ndb := db.Model(&userInfo).Update("cur_school_name", schoolName)
	return ndb.Error
}

func SetUserInfo(userInfo *model.UserInfo) error {
	if userInfo == nil {
		return nil
	}
	oldUserInfo, err := GetUserInfo(userInfo.OpenId)
	if err != nil {
		return err
	}
	if oldUserInfo != nil {
		strJson, err := jsoniter.MarshalToString(userInfo)
		if err != nil {
			return err
		}
		var userInfoMap map[string]interface{}
		err = jsoniter.UnmarshalFromString(strJson, &userInfoMap)
		if err != nil {
			return err
		}
		delete(userInfoMap, "open_id")
		return db.Model(&model.UserInfo{}).LogMode(true).Where("open_id = ?", userInfo.OpenId).Updates(userInfoMap).Error
	}
	return db.Model(&model.UserInfo{}).Save(userInfo).Error
}
