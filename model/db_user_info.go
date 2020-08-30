package model

type UserInfo struct {
	OpenId        string   `json:"open_id",grom:"open_id"`
	CurSchoolName *string  `json:"cur_school_name",grom:"cur_school_name"`
	PhoneNum      *string  `json:"phone_num",grom:"phone_num"`
	WxNum         *string  `json:"wx_num",grom:"wx_num"`
	Name          *string  `json:"name",grom:"name"`
	Sex           *int     `json:"sex",grom:"sex"`
	Extra         *string  `json:"extra"grom:"extra"`
	WxName        *string  `json:"wx_name",grom:"wx_name"`             //微信名；可选
	Latitude      *float64 `json:"latitude",grom:"latitude"`           //纬度；可选
	Longitude     *float64 `json:"longitude",grom:"longitude"`         //经度；可选
	ProfilePhoto  *string  `json:"profile_photo",grom:"profile_photo"` //头像；可选
}

func (s *UserInfo) TableName() string {
	return "user"
}
