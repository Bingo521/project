package model

type Image struct {
	Uri string `json:"uri",gorm:"uri"`
	Type string `json:"type",gorm:"type"`
}

func (img *Image)TableName()string{
	return "image_uri"
}
