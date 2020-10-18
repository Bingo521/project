package model

type UserInfo struct {
	OpenId        string   `json:"open_id",grom:"open_id"`
	CurSchoolName *string  `json:"cur_school_name,omitempty",grom:"cur_school_name"`
	PhoneNum      *string  `json:"phone_num,omitempty",grom:"phone_num"`
	WxNum         *string  `json:"wx_num,omitempty",grom:"wx_num"`
	Name          *string  `json:"name,omitempty",grom:"name"`
	Sex           *int     `json:"sex,omitempty",grom:"sex"`
	Extra         *string  `json:"extra,omitempty"grom:"extra"`
	WxName        *string  `json:"wx_name,omitempty",grom:"wx_name"`             //微信名；可选
	Latitude      *float64 `json:"latitude,omitempty",grom:"latitude"`           //纬度；可选
	Longitude     *float64 `json:"longitude,omitempty",grom:"longitude"`         //经度；可选
	ProfilePhoto  *string  `json:"profile_photo,omitempty",grom:"profile_photo"` //头像；可选
}

func (s *UserInfo) TableName() string {
	return "user"
}
