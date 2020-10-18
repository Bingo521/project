package db

import (
	"encoding/json"
	"my_project/model"
	"time"
)

func CreateSecondHand(openId string, message_id int64, content string, uris []string, price float32, category string) (*model.SecondHand, error) {
	message := model.SecondHand{}
	message.OpenId = openId
	message.MessageId = message_id
	message.CreateTime = time.Now()
	message.ModifyTime = message.CreateTime
	message.Content = content
	imageUris, err := json.Marshal(uris)
	if err != nil {
		return nil, err
	}
	message.ImageUris = string(imageUris)
	err = db.Save(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func GetSecondHandByOpenId(openId string, index int64, count int64) ([]model.SecondHand, error) {
	messages := make([]model.SecondHand, count)
	err := db.LogMode(true).Where("open_id = ?", openId).Order("create_time desc").Limit(count).Find(&messages).Error
	if err != nil {
		return []model.SecondHand{}, err
	}
	return messages, nil
}

func GetSecondHandTimeLine(firstTime int64, index int32, count int32) ([]model.SecondHand, error) {
	messages := make([]model.SecondHand, count)
	createTime := time.Unix(firstTime, 0).Format("2006-01-02 15:04:05")
	err := db.LogMode(true).LogMode(true).Where("create_time <= ?", createTime).Order("create_time desc").Offset(index).Limit(count).Find(&messages).Error
	if err != nil {
		return []model.SecondHand{}, err
	}
	return messages, nil
}

func GetSecondHandTimeLineByCategory(firstTime int64, index int32, count int32, category string) ([]model.SecondHand, error) {
	messages := make([]model.SecondHand, count)
	createTime := time.Unix(firstTime, 0).Format("2006-01-02 15:04:05")
	err := db.LogMode(true).LogMode(true).Where("create_time <= ? and category = ?", createTime, category).Order("create_time desc").Offset(index).Limit(count).Find(&messages).Error
	if err != nil {
		return []model.SecondHand{}, err
	}
	return messages, nil
}

func UpdateSecondHand(messageID int64, params map[string]interface{}) error {
	return db.Model(&model.SecondHand{}).Where("message_id = ?", messageID).Updates(params).Error
}
