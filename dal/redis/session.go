package redis

import (
	"encoding/json"
	"my_project/logs"
	"my_project/model"
	"time"
)

func SetSession(sessionInfo *model.WxLoginMainInfo) error {
	if sessionInfo == nil {
		return nil
	}
	logs.Info("set session key=%v", sessionInfo.SessionKey)
	val, err := json.Marshal(sessionInfo)
	if err != nil {
		return err
	}
	return Set(sessionInfo.SessionKey, string(val))
}

func GetSession(key string) (*model.WxLoginMainInfo, error) {
	logs.Info("get session key=%v", key)
	resp, err := Get(key)
	if err != nil {
		return nil, err
	}
	sessionInfo := model.WxLoginMainInfo{}
	err = json.Unmarshal([]byte(resp), &sessionInfo)
	if err != nil {
		return nil, err
	}
	return &sessionInfo, nil
}

func FlushSession(sessionKey string) error {
	return SetTimeOut(sessionKey, time.Second*3600*24)
}
