package model

import "time"

type Message struct {
	MessageId    int64     `json:"message_id",gorm:"message_id"`
	OpenId       string    `json:"open_id",gorm:"open_id"`
	Type         int32     `json:"type",gorm:"type"`
	Content      string    `json:"content",grom:"content"`
	ImageUris    string    `json:"image_uris",gorm:"image_uris"`
	CreateTime   time.Time `json:"create_time",gorm:"create_time"`
	ModifyTime   time.Time `json:"modify_time",gorm:"modify_time"`
	DiggCount    int32     `json:"digg_count",gorm:"modify_time"`
	CommentCount int32     `json:"comment_count",gorm:"comment_count"`
}

func (m *Message) TableName() string {
	return "message"
}
