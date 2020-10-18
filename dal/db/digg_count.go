package db

import (
	"github.com/jinzhu/gorm"
	"my_project/model"
)

func GetDiggCount(messageID int64) (*model.DiggCount, error) {
	diggInfo := model.DiggCount{}
	if err := db.Model(&model.DiggCount{}).Where("message_id = ?", messageID).First(&diggInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &diggInfo, nil
}

func SetDiggCount(messageID int64, count int64) error {
	diggInfo, err := GetDiggCount(messageID)
	if err != nil {
		return err
	}
	if diggInfo == nil {
		return db.Save(&model.DiggCount{MessageId: messageID, Count: count}).Error
	}
	return db.Model(&model.DiggCount{}).Where("message_id = ?", messageID).Updates(map[string]interface{}{"count": count}).Error
}

func MGetDiggCount(messageIDs []int64) (map[int64]int64, error) {
	diggInfos := make([]model.DiggCount, len(messageIDs))
	if err := db.Model(&model.DiggCount{}).Where("message_id in (?)", messageIDs).Find(&diggInfos).Error; err != nil {
		return nil, err
	}
	messageId2DiggCount := make(map[int64]int64, len(messageIDs))
	for _, diggInfo := range diggInfos {
		messageId2DiggCount[diggInfo.MessageId] = diggInfo.Count
	}
	return messageId2DiggCount, nil
}
