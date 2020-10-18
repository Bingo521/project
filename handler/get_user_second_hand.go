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
	"my_project/proto_gen/second_hand"
	"my_project/util"
	"strconv"
)

type GetSecondHandHandler struct {
	c    *gin.Context
	comm *model.CtxComm
	req  *second_hand.GetUserSecondHandRequest

	messages      []model.SecondHand
	userInfo      map[string]*model.UserInfo
	mid2DiggCount map[int64]int64
	hasMore       bool
}

func NewGetSecondHandHandler(c *gin.Context) *GetSecondHandHandler {
	comm := util.GetCtxComm(c)
	return &GetSecondHandHandler{
		c:    c,
		comm: comm,
	}
}

func (h *GetSecondHandHandler) getReq() (*second_hand.GetUserSecondHandRequest, error) {
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
	return &second_hand.GetUserSecondHandRequest{
		Index:  int32(index),
		Count:  int32(count),
		OpenId: openId,
	}, nil
}

func (h *GetSecondHandHandler) Execute() *second_hand.GetUserSecondHandResonse {
	req, err := h.getReq()
	if err != nil {
		logs.Warn("[GetSecondHandHandler] make req err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.req = req
	if err := h.check(); err != nil {
		logs.Warn("[GetSecondHandHandler] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.loadData(); err != nil {
		logs.Warn("[GetSecondHandHandler] loadData err:%v", err)
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

func (h *GetSecondHandHandler) check() error {
	if h.req == nil {
		return fmt.Errorf("req is nil")
	}
	if h.req.Index < 0 || h.req.Count <= 0 || h.req.Count > 20 {
		return fmt.Errorf("param is illegal")
	}
	return nil
}

func (h *GetSecondHandHandler) makeErrResp(errCode int32, errMessage string) *second_hand.GetUserSecondHandResonse {
	return &second_hand.GetUserSecondHandResonse{
		StatusCode: errCode,
		Message:    errMessage,
	}
}

func (h *GetSecondHandHandler) transToClientMessage() []*second_hand.SecondHandInfo {
	messageInfos := make([]*second_hand.SecondHandInfo, 0, len(h.messages))
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
		messageInfo := &second_hand.SecondHandInfo{
			MessageId:  messageItem.MessageId,
			Content:    messageItem.Content,
			Urls:       uris,
			CreateTime: messageItem.CreateTime.Unix(),
		}
		clientUserInfo := &second_hand.SecondHandUserInfo{
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

func (h *GetSecondHandHandler) GetOpenIDs(messages []model.SecondHand) []string {
	opendIds := make([]string, 0, len(messages))
	for _, mess := range messages {
		opendIds = append(opendIds, mess.OpenId)
	}
	return opendIds
}

func (h *GetSecondHandHandler) loadUserInfo() error {
	openIDs := h.GetOpenIDs(h.messages)
	userInfo, err := db.MGetUserInfo(openIDs)
	if err != nil {
		return err
	}
	h.userInfo = userInfo
	return nil
}

func (h *GetSecondHandHandler) loadMessages() error {
	messages, err := db.GetSecondHandByOpenId(h.comm.OpenId, int64(h.req.Index), int64(h.req.Count+1))
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

func (h *GetSecondHandHandler) loadData() error {
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

func (h *GetSecondHandHandler) getMessageIds(messages []model.SecondHand) []int64 {
	messageIds := make([]int64, 0, len(messages))
	for _, mess := range messages {
		messageIds = append(messageIds, mess.MessageId)
	}
	return messageIds
}

func (h *GetSecondHandHandler) loadDiggCount() error {
	messageIDs := h.getMessageIds(h.messages)
	diggCount, err := db.MGetDiggCount(messageIDs)
	if err != nil {
		return err
	}
	h.mid2DiggCount = diggCount
	return nil
}
