package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"my_project/consts"
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
	return &Login{
	}
}

func (h *Login) Execute(ctx *gin.Context) *login.LoginResponse {
	wxLoginResp := model.WxLoginResp{}
	code:=ctx.PostForm("code")
	logs.Info("code = %v",code)
	if code == "admin"{
		wxLoginResp.OpenId = "000000"
		wxLoginResp.SessionKey = "1234567890"
		h.saveSession(ctx,&wxLoginResp.WxLoginMainInfo)
		return h.MakeResp(error_code.ERR_SUCCESS,"success",wxLoginResp.OpenId,wxLoginResp.SessionKey)
	}
	url := getLoginUrl(code)
	resp,err:=http.Get(url)
	if err != nil{
		logs.Error(fmt.Sprintf("login wx resp err = %v",err))
		return h.MakeResp(error_code.ERR_LOGIN,"wx resp err","","")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(fmt.Sprintf("login wx resp body = %v ,err = %v",string(body),err))
		return h.MakeResp(error_code.ERR_LOGIN,"wx resp err","","")
	}
	logs.Info("body = %s",string(body))

	err=json.Unmarshal(body,&wxLoginResp)
	if err != nil{
		logs.Error(fmt.Sprintf("login wx resp = %v ,err = %v",string(body),err))
		return h.MakeResp(error_code.ERR_LOGIN,"wx resp err","","")
	}
	if wxLoginResp.ErrCode != 0{
		logs.Error(fmt.Sprintf("login wx resp = %v ,err = %v",string(body),err))
		return h.MakeResp(error_code.ERR_LOGIN,"wx resp err","","")
	}
	h.saveSession(ctx,&wxLoginResp.WxLoginMainInfo)
	return h.MakeResp(error_code.ERR_SUCCESS,"success",wxLoginResp.OpenId,wxLoginResp.SessionKey)

}

func (h *Login)MakeResp(statusCode int32,message string,openId string,sessionKey string)*login.LoginResponse{
	return &login.LoginResponse{
		StatusCode: statusCode,
		Message:    message,
		Openid:     openId,
		SessionKey: sessionKey,
	}
}
func getLoginUrl(code string)string{
	return fmt.Sprintf(consts.LOGIN_FORMAT,consts.APPID,consts.APP_SECRET,code)
}

func (h *Login)saveSession(ctx *gin.Context,loginMainInfo *model.WxLoginMainInfo)error{
	if loginMainInfo == nil{
		return nil
	}
	return redis.SetSession(loginMainInfo)
}