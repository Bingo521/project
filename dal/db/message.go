package db

import (
	"encoding/json"
	"my_project/model"
	"time"
)

func CreateMessage(openId string,message_id int64,messageType int32,content string,uris []string)(*model.Message,error){
	message := model.Message{}
	message.OpenId = openId
	message.MessageId = message_id
	message.Type = messageType
	message.CreateTime = time.Now()
	message.ModifyTime = message.CreateTime
	message.Message = content
	imageUris,err := json.Marshal(uris)
	if err != nil{
		return nil, err
	}
	message.ImageUris = string(imageUris)
	err = db.Save(&message).Error
	if err != nil{
		return nil,err
	}
	return &message,nil
}

func GetMessageByOpenId(openId string,index int64,count int64)([]model.Message,error){
	messages := make([]model.Message,count)
	err:=db.LogMode(true).Where("open_id = ?",openId).Order("create_time desc").Limit(count).Find(&messages).Error
	if err != nil{
		return []model.Message{},err
	}
	return messages,nil
}
