package db

import (
	"github.com/jinzhu/gorm"
	"my_project/model"
)

const (
	DIGG_TYPE_DIGG   = 1
	DIGG_TYPE_CANCEL = 2
)

func Digg(openID string, messageID int64) error {
	diggInfo, err := GetDiggStatus(openID, messageID)
	if err != nil {
		return err
	}
	if diggInfo == nil {
		diggInfo = &model.Digg{}
		diggInfo.MessageID = messageID
		diggInfo.OpenID = openID
		return db.Save(diggInfo).Error
	}
	return nil
}

func DiggCancel(openID string, messageID int64) error {
	diggInfo, err := GetDiggStatus(openID, messageID)
	if err != nil {
		return err
	}
	if diggInfo != nil {
		return db.Model(&model.Digg{}).Where("open_id = ? and message_id = ?", openID, messageID).Delete(model.Digg{}).Error
	}
	return nil
}

func GetDiggStatus(openID string, messageID int64) (*model.Digg, error) {
	diggInfo := model.Digg{}
	if err := db.Model(&model.Digg{}).Where("open_id = ? and message_id = ?", openID, messageID).First(&diggInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &diggInfo, nil
}

func GetDiggCountByMessageId(messageID int64) (int64, error) {
	var count int64
	if err := db.Model(&model.Digg{}).Where(" message_id = ?", messageID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
