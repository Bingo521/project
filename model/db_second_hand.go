package model

import "time"

type SecondHand struct {
	MessageId    int64     `json:"message_id",gorm:"message_id"`
	OpenId       string    `json:"open_id",gorm:"open_id"`
	Content      string    `json:"content",grom:"content"`
	ImageUris    string    `json:"image_uris",gorm:"image_uris"`
	CreateTime   time.Time `json:"create_time",gorm:"create_time"`
	ModifyTime   time.Time `json:"modify_time",gorm:"modify_time"`
	Money        float32   `json:"money",gorm:"money"`
	Category     string    `json:"category",gorm:"category"`
	Extra        string    `json:"extra",gorm:"extra"`
	DiggCount    int32     `json:"digg_count",gorm:"digg_count"`
	CommentCount int32     `json:"comment_count",gorm:"comment_count"`
}

func (m *SecondHand) TableName() string {
	return "message"
}
