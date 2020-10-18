package model

type Digg struct {
	OpenID    string `json:"open_id",gorm:"open_id"`
	MessageID int64  `json:"message_id",gorm:"message_id"`
}

func (d *Digg) TableName() string {
	return "digg"
}
