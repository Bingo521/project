package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/message"
	"my_project/util"
	"strconv"
)

type GetUserMessageHandler struct {
	c    *gin.Context
	comm *model.CtxComm
	req  *message.GetUserMessageRequest

	messages      []model.Message
	userInfo      map[string]*model.UserInfo
	mid2DiggCount map[int64]int64
	hasMore       bool
}

func NewGetUserMessageHandler(c *gin.Context) *GetUserMessageHandler {
	comm := util.GetCtxComm(c)
	return &GetUserMessageHandler{
		c:    c,
		comm: comm,
	}
}

func (h *GetUserMessageHandler) getReq() (*message.GetUserMessageRequest, error) {
	if h.c == nil {
		return nil, errors.New("context is nil")
	}
	strIndex := h.c.Query("index")
	strCount := h.c.Query("count")
	openId := h.c.Query("open_id")
	logs.Info("index = %v,count = %v,open_id = %v", strIndex, strCount, openId)
	if strIndex == "" || strCount == "" || openId == "" {
		return nil, fmt.Errorf("param not find")
	}
	index, err := strconv.ParseInt(strIndex, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("index is err:%v", err)
	}
	count, err := strconv.ParseInt(strCount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("count is err:%v", err)
	}
	if count == 0 {
		return nil, nil
	}
	return &message.GetUserMessageRequest{
		Index:  int32(index),
		Count:  int32(count),
		OpenId: openId,
	}, nil
}

func (h *GetUserMessageHandler) Execute() *message.GetUserMessageResonse {
	req, err := h.getReq()
	if err != nil {
		logs.Warn("[GetUserMessageHandler] make req err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.req = req
	if err := h.check(); err != nil {
		logs.Warn("[GetUserMessageHandler] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.loadData(); err != nil {
		logs.Warn("[GetUserMessageHandler] loadData err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	respMess := h.transToClientMessage()
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.HasMore = h.hasMore
	resp.MessageInfos = respMess
	resp.NextIndex = req.Index + int32(len(h.messages))
	resp.Count = int64(len(respMess))
	return resp
}

func (h *GetUserMessageHandler) check() error {
	if h.req == nil {
		return fmt.Errorf("req is nil")
	}
	if h.req.Index < 0 || h.req.Count <= 0 || h.req.Count > 20 {
		return fmt.Errorf("param is illegal")
	}
	return nil
}

func (h *GetUserMessageHandler) makeErrResp(errCode int32, errMessage string) *message.GetUserMessageResonse {
	return &message.GetUserMessageResonse{
		StatusCode: errCode,
		Message:    errMessage,
	}
}

func (h *GetUserMessageHandler) transToClientMessage() []*message.MessageInfo {
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
			Content:     messageItem.Content,
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
		messageInfo.DiggCount = int32(h.mid2DiggCount[messageItem.MessageId])
		messageInfos = append(messageInfos, messageInfo)
	}
	return messageInfos
}

func (h *GetUserMessageHandler) GetOpenIDs(messages []model.Message) []string {
	opendIds := make([]string, 0, len(messages))
	for _, mess := range messages {
		opendIds = append(opendIds, mess.OpenId)
	}
	return opendIds
}

func (h *GetUserMessageHandler) loadUserInfo() error {
	openIDs := h.GetOpenIDs(h.messages)
	userInfo, err := db.MGetUserInfo(openIDs)
	if err != nil {
		return err
	}
	h.userInfo = userInfo
	return nil
}

func (h *GetUserMessageHandler) loadMessages() error {
	messages, err := db.GetMessageByOpenId(h.req.OpenId, int64(h.req.Index), int64(h.req.Count+1))
	if err != nil {
		return err
	}
	needMessage := messages
	hasMore := false
	if len(messages) > int(h.req.Count) {
		needMessage = needMessage[:h.req.Count]
		hasMore = true
	}
	h.messages = needMessage
	h.hasMore = hasMore
	return nil
}

func (h *GetUserMessageHandler) loadData() error {
	wg := util.WaitGroup{}
	if err := wg.Go(func() error {
		if err := h.loadMessages(); err != nil {
			logs.Warn("[GetMessageByTimeLine] GetUserInfo err:%v", err)
			return err
		}
		return nil
	}).Wait(); err != nil {
		return err
	}
	wg.Go(func() error {
		if err := h.loadUserInfo(); err != nil {
			logs.Warn("[GetMessageByTimeLine] GetMessageTimeLine err:%v", err)
			return err
		}
		return nil
	}).Go(func() error {
		if err := h.loadDiggCount(); err != nil {
			logs.Warn("[GetMessageByTimeLine] ")
			return err
		}
		return nil
	}).Wait()
	return nil
}

func (h *GetUserMessageHandler) getMessageIds(messages []model.Message) []int64 {
	messageIds := make([]int64, 0, len(messages))
	for _, mess := range messages {
		messageIds = append(messageIds, mess.MessageId)
	}
	return messageIds
}

func (h *GetUserMessageHandler) loadDiggCount() error {
	messageIDs := h.getMessageIds(h.messages)
	diggCount, err := db.MGetDiggCount(messageIDs)
	if err != nil {
		return err
	}
	h.mid2DiggCount = diggCount
	return nil
}
