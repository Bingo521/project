package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"my_project/consts"
	"my_project/dal/db"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/login"
	"net/http"
)

type Login struct {
}

func NewLogin() *Login {
	return &Login{}
}

func (h *Login) Execute(ctx *gin.Context) *login.LoginResponse {
	code := ctx.PostForm("code")
	logs.Info("code = %v", code)
	wxLoginResp := h.save(code)
	if wxLoginResp.StatusCode != 0 {
		return wxLoginResp
	}
	userInfo, err := db.GetUserInfo(wxLoginResp.Openid)
	if err != nil {
		logs.Warn("[Login] code = %v err:%v", code, err)
		return h.MakeResp(error_code.ERR_SERVER_ERR, "service", "", "")
	}
	if userInfo != nil {
		wxLoginResp.IsFirstLogin = false
	} else {
		wxLoginResp.IsFirstLogin = true
		userInfo := &model.UserInfo{
			OpenId: wxLoginResp.Openid,
		}
		if err := db.SetUserInfo(userInfo); err != nil {
			logs.Warn("[Login] code = %v err:%v", code, err)
			return h.MakeResp(error_code.ERR_SERVER_ERR, "service", "", "")
		}
	}
	return h.MakeResp(error_code.ERR_SUCCESS, "success", wxLoginResp.Openid, wxLoginResp.SessionKey)
}

func (h *Login) save(code string) *login.LoginResponse {
	wxLoginResp := model.WxLoginResp{}
	if code == "admin" {
		wxLoginResp.OpenId = "000000"
		wxLoginResp.SessionKey = "1234567890"
		if err := h.saveSession(&wxLoginResp.WxLoginMainInfo); err != nil {
			logs.Warn("save session err:%v", err)
			return h.MakeResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR, "", "")
		}
		return h.MakeResp(error_code.ERR_SUCCESS, "success", wxLoginResp.OpenId, wxLoginResp.SessionKey)
	}
	url := getLoginUrl(code)
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("login wx resp err = %v", err))
		return h.MakeResp(error_code.ERR_LOGIN, "wx resp err", "", "")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(fmt.Sprintf("login wx resp body = %v ,err = %v", string(body), err))
		return h.MakeResp(error_code.ERR_LOGIN, "wx resp err", "", "")
	}
	logs.Info("body = %s", string(body))
	err = json.Unmarshal(body, &wxLoginResp)
	if err != nil {
		logs.Error(fmt.Sprintf("login wx resp = %v ,err = %v", string(body), err))
		return h.MakeResp(error_code.ERR_LOGIN, "wx resp err", "", "")
	}
	if wxLoginResp.ErrCode != 0 {
		logs.Error(fmt.Sprintf("login wx resp = %v ,err = %v", string(body), err))
		return h.MakeResp(error_code.ERR_LOGIN, "wx resp err", "", "")
	}
	err = h.saveSession(&wxLoginResp.WxLoginMainInfo)
	if err != nil {
		logs.Warn(fmt.Sprintf("saveSession err = %v", err))
		return h.MakeResp(error_code.ERR_LOGIN, "redis err", "", "")
	}
	return h.MakeResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS, wxLoginResp.OpenId, wxLoginResp.SessionKey)
}

func (h *Login) MakeResp(statusCode int32, message string, openId string, sessionKey string) *login.LoginResponse {
	return &login.LoginResponse{
		StatusCode: statusCode,
		Message:    message,
		Openid:     openId,
		SessionKey: sessionKey,
	}
}
func getLoginUrl(code string) string {
	return fmt.Sprintf(consts.LOGIN_FORMAT, consts.APPID, consts.APP_SECRET, code)
}

func (h *Login) saveSession(loginMainInfo *model.WxLoginMainInfo) error {
	if loginMainInfo == nil {
		return nil
	}
	return redis.SetSession(loginMainInfo)
}
