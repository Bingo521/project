package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"my_project/consts"
	"my_project/dal/db"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/message"
	"my_project/util"
	"strconv"
)

type MessageHandler struct {
	c    *gin.Context
	comm *model.CtxComm
	req  *message.CreateMessageRequest
}

func NewMessageHandler(c *gin.Context) *MessageHandler {
	comm := util.GetCtxComm(c)
	return &MessageHandler{
		c:    c,
		comm: comm,
	}
}

func (h *MessageHandler) Execute() *message.CreateMessageResponse {
	resp := message.CreateMessageResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	if h.c == nil || h.comm == nil {
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	var err error
	openId := h.comm.OpenId
	h.req, err = h.getReq()
	if err != nil {
		logs.Error("get req err:%v", err)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	userInfos, err := db.MGetUserInfo([]string{openId})
	if err != nil {
		logs.Warn("[MessageHandler] openId = %v getUserInfo err:%v", openId, err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	userInfo, find := userInfos[openId]
	if !find {
		logs.Warn("[MessageHandler] openId = %v getUserInfo err:not find", openId)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	mId, err := redis.Inc("message_id")
	if err != nil {
		logs.Error("get message ID err:%v", err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	mess, err := db.CreateMessage(openId, mId, consts.MESSAGE_TYPE_NORMAL_THREAD, h.req.Content, h.req.Uris)
	if err != nil {
		logs.Error("create message err :%v", err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	resp.MessageInfo = &message.MessageInfo{
		Content:    h.req.Content,
		Urls:       h.req.Uris,
		MessageId:  mId,
		CreateTime: mess.CreateTime.Unix(),
		UserInfo:   transToLocalUserInfo(userInfo),
	}
	return &resp
}

func (h *MessageHandler) getReq() (*message.CreateMessageRequest, error) {
	content := h.c.PostForm("content")
	strUris := h.c.PostForm("uris")
	messageType := h.c.PostForm("message_type")
	logs.Info("[MessageHandler] content = %v,uris = %v,message_type = %v", content, strUris, messageType)
	var uris []string
	err := json.Unmarshal([]byte(strUris), &uris)
	if err != nil {
		logs.Error("get req err:uris = %v,err = %v", strUris, err)
		return nil, err
	}
	iMessageType, err := strconv.ParseInt(messageType, 10, 64)
	if err != nil {
		logs.Warn("[MessageHandler] message_type = %v is illegal", messageType)
		return nil, err
	}
	return &message.CreateMessageRequest{
		Content:     content,
		Uris:        uris,
		MessageType: int32(iMessageType),
	}, nil
}

func (h *MessageHandler) checkParams() error {
	if h.req == nil {
		return errors.New("param is nil")
	}
	if h.req.Content == "" || len(h.req.Uris) == 0 {
		return errors.New("param illegal")
	}
	return nil
}

func transToLocalUserInfo(userInfo *model.UserInfo) *message.MessageUserInfo {
	if userInfo == nil {
		return nil
	}
	localUserInfo := message.MessageUserInfo{}
	localUserInfo.OpenId = userInfo.OpenId
	if userInfo.ProfilePhoto != nil {
		localUserInfo.ProfilePhoto = *userInfo.ProfilePhoto
	}
	if userInfo.Sex != nil {
		localUserInfo.Sex = int32(*userInfo.Sex)
	}
	if userInfo.Name != nil {
		localUserInfo.UserName = *userInfo.Name
	}
	if localUserInfo.UserName == "" && userInfo.WxName != nil {
		localUserInfo.UserName = *userInfo.WxName
	}
	return &localUserInfo
}
