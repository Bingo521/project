package model

type ClassRecord struct {
	OpenId string `json:"open_id",gorm:"open_id"`
	SchoolName string `json:"school_name",gorm:"school_name"`
	Week int32 `json:"week",gorm:"week"`
	Class string `json:"class",gorm:"class"`
	ClassInfos [][]*ClassInfo `json:"-"`
}

type ClassInfo struct {
	ClassName string `json:"class_name"`
	TeacherName string `json:"teacher_name"`
}

func (c *ClassRecord)TableName()string{
	return "class_record"
}