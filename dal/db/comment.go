package db

import (
	"encoding/json"
	"my_project/model"
	"time"
)

func CreateComment(commentID int64, messageID int64, content string, uris []string, commentType int32) error {
	imageUris, err := json.Marshal(uris)
	if err != nil {
		return err
	}
	commentInfo := model.Comment{}
	commentInfo.CommentID = commentID
	commentInfo.MessageID = messageID
	commentInfo.DiggCount = 0
	commentInfo.Content = content
	commentInfo.Type = commentType
	commentInfo.CreateTime = time.Now()
	commentInfo.ModifyTime = commentInfo.CreateTime
	commentInfo.ImageUris = string(imageUris)
	return db.Model(&model.Comment{}).Save(&commentInfo).Error
}

func GetCommentByTimeLine(messageID int64, startTime int64, index, count int32) ([]model.Comment, error) {
	comments := make([]model.Comment, count)
	createTime := time.Unix(startTime, 0).Format("2006-01-02 15:04:05")
	err := db.LogMode(true).Where("message_id = ? and create_time <= ?", messageID, createTime).Order("create_time desc").Offset(index).Limit(count).Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}

func DeleteCommentByMessageId(messageID int64) error {
	return db.Model(&model.Comment{}).Where("comment_id = ?", messageID).Delete(model.Comment{}).Error
}
