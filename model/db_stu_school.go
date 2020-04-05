package model

type StuSchool struct {
	OpenId string `json:"open_id",gorm:"open_id"`
	SchoolName string `json:"school_name",gorm:"school_name"`
	StuId string `json:"stu_id",gorm:"stu_id"`
	StuPassword string `json:"stu_password",gorm:"stu_password"`
	FirstWeekDate int64 `json:"first_week_date",gorm:"first_week_date"`
	Extra string `json:"extra",gorm:"extra"`
}

func (s *StuSchool)TableName()string{
	return "stu_school"
}