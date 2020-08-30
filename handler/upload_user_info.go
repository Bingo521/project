package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/user_info"
	"my_project/util"
	"strconv"
)

type UploadUserInfo struct {
	c       *gin.Context
	comm    *model.CtxComm
	r       *user_info.UserInfoRequest
	dbModel *model.UserInfo
}

func NewUploadUserInfoHandler(c *gin.Context) *UploadUserInfo {
	return &UploadUserInfo{
		c: c,
	}
}

func (h *UploadUserInfo) Handle() *user_info.UserInfoResponse {
	userInfo, err := h.makeReq()
	if err != nil {
		logs.Warn("[UploadUserInfo] make req err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.dbModel = userInfo
	if err := h.check(); err != nil {
		logs.Warn("[UploadUserInfo] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := db.SetUserInfo(h.dbModel); err != nil {
		logs.Warn("[UploadUserInfo] setUserInfo err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	return h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
}

func (h *UploadUserInfo) makeErrResp(errCode int32, errMess string) *user_info.UserInfoResponse {
	return &user_info.UserInfoResponse{
		StatusCode: errCode,
		Message:    errMess,
	}
}

func (h *UploadUserInfo) makeReq() (*model.UserInfo, error) {
	h.comm = util.GetCtxComm(h.c)
	if h.comm == nil {
		return nil, errors.New("comm is err")
	}
	curSchoolName := h.c.PostForm("cur_school_name")
	phoneNum := h.c.PostForm("phone_num")
	wxNum := h.c.PostForm("wx_num")
	name := h.c.PostForm("name")
	sex := h.c.PostForm("sex")
	latitude := h.c.PostForm("latitude")
	longitude := h.c.PostForm("longitude")
	profilePhoto := h.c.PostForm("profile_photo")
	logs.Info("openId = %v,CurSchoolName = %v,PhoneNum = %v,WxNum = %v,Name = %v,sex = %v,latitude = %v,latitude = %v,longitude = %v,ProfilePhoto = %v",
		h.comm.OpenId, curSchoolName, phoneNum, wxNum, name, sex, latitude, longitude, profilePhoto)
	iSex, _ := strconv.ParseInt(sex, 10, 64)
	iLatitude, _ := strconv.ParseFloat(latitude, 64)
	iLongitude, _ := strconv.ParseFloat(longitude, 64)
	userInfo := model.UserInfo{}
	userInfo.OpenId = h.comm.OpenId
	if curSchoolName != "" {
		userInfo.CurSchoolName = &curSchoolName
	}
	if phoneNum != "" {
		userInfo.PhoneNum = &phoneNum
	}
	if wxNum != "" {
		userInfo.WxNum = &wxNum
	}
	if name != "" {
		userInfo.Name = &name
	}
	if sex != "" {
		iisex := int(iSex)
		userInfo.Sex = &iisex
	}
	if longitude != "" {
		userInfo.Longitude = &iLongitude
	}
	if latitude != "" {
		userInfo.Latitude = &iLatitude
	}
	if profilePhoto != "" {
		userInfo.ProfilePhoto = &profilePhoto
	}
	return &userInfo, nil
}

func (h *UploadUserInfo) check() error {
	if h.dbModel == nil {
		return errors.New("参数错误")
	}
	//TODO 检查其它参数
	return nil
}
