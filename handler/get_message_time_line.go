package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/message"
	"strconv"
)

type GetMessageByTimeLine struct {
	c        *gin.Context
	comm     *model.CtxComm
	r        *message.GetMessageRequest
	messages []model.Message
	userInfo map[string]*model.UserInfo
	hasMore  bool
}

func NewGetMessageByTimeLine(c *gin.Context) *GetMessageByTimeLine {
	return &GetMessageByTimeLine{
		c: c,
	}
}

func (h *GetMessageByTimeLine) Handle() *message.GetMessageResponse {
	req, err := h.makeReq()
	if err != nil {
		logs.Warn("[GetMessageByTimeLine] make req err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.r = req
	if err := h.check(); err != nil {
		logs.Warn("[GetMessageByTimeLine] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.loadMessages(); err != nil {
		logs.Warn("[GetMessageByTimeLine] GetUserInfo err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	if err := h.loadUserInfo(); err != nil {
		logs.Warn("[GetMessageByTimeLine] GetMessageTimeLine err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	respMess := h.transToClientMessage()
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.HasMore = h.hasMore
	resp.MessageInfos = respMess
	resp.NextIndex = int32(req.Index + int32(len(h.messages)))
	return resp
}

func (h *GetMessageByTimeLine) makeReq() (*message.GetMessageRequest, error) {
	index := h.c.Query("index")
	iindex, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("index = %v is illegal", index)
	}
	count := h.c.Query("count")
	icount, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("count = %v is illegal", count)
	}
	firstTime := h.c.Query("first_time")
	iFirstTime, err := strconv.ParseInt(firstTime, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("first_time = %v is illegal", firstTime)
	}
	return &message.GetMessageRequest{
		Index:     int32(iindex),
		Count:     int32(icount),
		FirstTime: iFirstTime,
	}, nil
}

func (h *GetMessageByTimeLine) check() error {
	if h.r == nil {
		return fmt.Errorf("req is nil")
	}
	if h.r.Index < 0 || h.r.Count <= 0 || h.r.Count > 20 {
		return fmt.Errorf("param is illegal")
	}
	return nil
}

func (h *GetMessageByTimeLine) makeErrResp(errCode int32, errMessage string) *message.GetMessageResponse {
	return &message.GetMessageResponse{
		StatusCode: errCode,
		Message:    errMessage,
	}
}

func (h *GetMessageByTimeLine) transToClientMessage() []*message.MessageInfo {
	messageInfos := make([]*message.MessageInfo, 0, len(h.messages))
	for _, messageItem := range h.messages {
		var uris []string
		if err := json.Unmarshal([]byte(messageItem.ImageUris), &uris); err != nil {
			logs.Error("message_id = %v image_uris = %v", messageItem.MessageId, messageItem.ImageUris)
			continue
		}
		userInfo, find := h.userInfo[messageItem.OpenId]
		if !find {
			logs.Warn("[GetMessageByTimeLine] transToClientMessage openId = %v not find UserInfo", messageItem.OpenId)
			continue
		}
		messageInfo := &message.MessageInfo{
			MessageId:   messageItem.MessageId,
			Content:     messageItem.Message,
			Urls:        uris,
			MessageType: messageItem.Type,
			CreateTime:  messageItem.CreateTime.Unix(),
		}
		clientUserInfo := &message.MessageUserInfo{
			OpenId: userInfo.OpenId,
		}
		if userInfo.ProfilePhoto != nil {
			clientUserInfo.ProfilePhoto = *userInfo.ProfilePhoto
		}
		if userInfo.Name != nil {
			clientUserInfo.UserName = *userInfo.Name
		}
		if userInfo.Sex != nil {
			clientUserInfo.Sex = int32(*userInfo.Sex)
		}
		messageInfo.UserInfo = clientUserInfo
		messageInfos = append(messageInfos, messageInfo)
	}
	return messageInfos
}

func (h *GetMessageByTimeLine) GetOpenIDs(messages []model.Message) []string {
	opendIds := make([]string, 0, len(messages))
	for _, mess := range messages {
		opendIds = append(opendIds, mess.OpenId)
	}
	return opendIds
}

func (h *GetMessageByTimeLine) loadUserInfo() error {
	openIDs := h.GetOpenIDs(h.messages)
	userInfo, err := db.MGetUserInfo(openIDs)
	if err != nil {
		return err
	}
	h.userInfo = userInfo
	return nil
}

func (h *GetMessageByTimeLine) loadMessages() error {
	messages, err := db.GetMessageTimeLine(h.r.FirstTime, h.r.Index, h.r.Count+1)
	if err != nil {
		return err
	}
	needMessage := messages
	hasMore := false
	if len(messages) > int(h.r.Count) {
		needMessage = needMessage[:h.r.Count]
		hasMore = true
	}
	h.messages = needMessage
	h.hasMore = hasMore
	return nil
}
