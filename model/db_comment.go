package model

import "time"

type Comment struct {
	CommentID  int64     `json:"comment_id",gorm:"comment_id"`
	MessageID  int64     `json:"message_id",gorm:"message_id"`
	Type       int32     `json:"type",gorm:"type"`
	Content    string    `json:"content",gorm:"content"`
	ImageUris  string    `json:"image_uris",gorm:"image_uris"`
	CreateTime time.Time `json:"create_time",gorm:"create_time"`
	ModifyTime time.Time `json:"modify_time",gorm:"modify_time"`
	Extra      string    `json:"extra",gorm:"extra"`
	DiggCount  int64     `json:"digg_count",gorm:"digg_count"`
	ReplyCount int64     `json:"reply_count",gorm:"reply_count"`
}

func (c *Comment) TableName() string {
	return "comment"
}

const (
	COMMENT_TYPE_COMMENT = 1
	COMMENT_TYPE_REPLY   = 2
)
