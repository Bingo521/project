package model

type DiggCount struct {
	MessageId int64 `json:"message_id",gorm:"message_id"`
	Count     int64 `json:"count",gorm:"count"`
}

func (d *DiggCount) TableName() string {
	return "digg_count"
}
