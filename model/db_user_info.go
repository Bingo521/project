package model

type UserInfo struct {
	OpenId string `json:"open_id",grom:"open_id"`
	CurSchoolName string `json:"cur_school_name",grom:"cur_school_name"`
	PhoneNum string `json:"phone_num",grom:"phone_num"`
	WxNum string `json:"wx_num",grom:"wx_num"`
	Name string `json:"name",grom:"name"`
	Sex int `json:"sex",grom:"sex"`
	Extra string  `json:"extra"grom:"extra"`
}

func (s *UserInfo)TableName()string{
	return "user"
}