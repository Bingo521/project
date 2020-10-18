package db

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"my_project/model"
	"time"
)

func CreateMessage(openId string, message_id int64, messageType int32, content string, uris []string) (*model.Message, error) {
	message := model.Message{}
	message.OpenId = openId
	message.MessageId = message_id
	message.Type = messageType
	message.CreateTime = time.Now()
	message.ModifyTime = message.CreateTime
	message.Content = content
	message.DiggCount = 0
	message.CommentCount = 0
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

func UpdateMessage(messageID int64, params map[string]interface{}) error {
	return db.Model(&model.Message{}).Where("message_id = ?", messageID).Updates(params).Error
}

func GetMessageByOpenId(openId string, index int64, count int64) ([]model.Message, error) {
	messages := make([]model.Message, count)
	err := db.LogMode(true).Where("open_id = ?", openId).Order("create_time desc").Limit(count).Find(&messages).Error
	if err != nil {
		return []model.Message{}, err
	}
	return messages, nil
}

func GetMessageTimeLine(firstTime int64, index int32, count int32) ([]model.Message, error) {
	messages := make([]model.Message, count)
	createTime := time.Unix(firstTime, 0).Format("2006-01-02 15:04:05")
	err := db.LogMode(true).Where("create_time <= ?", createTime).Order("create_time desc").Offset(index).Limit(count).Find(&messages).Error
	if err != nil {
		return []model.Message{}, err
	}
	return messages, nil
}

func GetMessageByMessageId(messageID int64) (*model.Message, error) {
	messageInfo := model.Message{}
	if err := db.Model(&model.Message{}).Where("message_id = ?", messageID).First(&messageInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &messageInfo, nil
}

func DeleteMessageByMessageId(messageID int64) error {
	return db.Model(&model.Message{}).Where("message_id = ?", messageID).Delete(model.Message{}).Error
}
