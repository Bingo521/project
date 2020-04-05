package model

type WxLoginResp struct {
	WxLoginMainInfo
	ErrCode	int32`json:"errcode"`
	ErrMsg	string	`json:"errmsg"`
}

type WxLoginMainInfo struct {
	OpenId	string	`json:"openid"`
	SessionKey	string `json:"session_key"`
	UnionId	string	`json:"unionid"`
}
