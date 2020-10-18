package handler

import (
	"encoding/json"
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

type GetSecondHandByTimeLine struct {
	c             *gin.Context
	comm          *model.CtxComm
	r             *second_hand.GetMessageRequest
	messages      []model.SecondHand
	userInfo      map[string]*model.UserInfo
	mid2DiggCount map[int64]int64
	hasMore       bool
}

func NewGetSecondHandByTimeLine(c *gin.Context) *GetSecondHandByTimeLine {
	return &GetSecondHandByTimeLine{
		c: c,
	}
}

func (h *GetSecondHandByTimeLine) Handle() *second_hand.GetMessageResponse {
	req, err := h.makeReq()
	if err != nil {
		logs.Warn("[GetSecondHandByTimeLine] make req err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.r = req
	if err := h.check(); err != nil {
		logs.Warn("[GetSecondHandByTimeLine] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.loadData(); err != nil {
		logs.Warn("[GetSecondHandByTimeLine] loadData err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	respMess := h.transToClientMessage()
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.HasMore = h.hasMore
	resp.MessageInfos = respMess
	resp.NextIndex = int32(req.Index + int32(len(h.messages)))
	resp.FirstTime = req.FirstTime
	return resp
}

func (h *GetSecondHandByTimeLine) makeReq() (*second_hand.GetMessageRequest, error) {
	index := h.c.Query("index")
	count := h.c.Query("count")
	firstTime := h.c.Query("first_time")
	category := h.c.Query("category")
	logs.Info("[GetSecondHandByTimeLine] index = %v,count = %v,firstTime = %v,category = %v", index, count, firstTime, category)
	iindex, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("index = %v is illegal", index)
	}
	icount, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("count = %v is illegal", count)
	}
	iFirstTime, err := strconv.ParseInt(firstTime, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("first_time = %v is illegal", firstTime)
	}

	return &second_hand.GetMessageRequest{
		Index:     int32(iindex),
		Count:     int32(icount),
		FirstTime: iFirstTime,
		Category:  category,
	}, nil
}

func (h *GetSecondHandByTimeLine) check() error {
	if h.r == nil {
		return fmt.Errorf("req is nil")
	}
	if h.r.Index < 0 || h.r.Count <= 0 || h.r.Count > 20 {
		return fmt.Errorf("param is illegal")
	}
	return nil
}

func (h *GetSecondHandByTimeLine) makeErrResp(errCode int32, errMessage string) *second_hand.GetMessageResponse {
	return &second_hand.GetMessageResponse{
		StatusCode: errCode,
		Message:    errMessage,
	}
}

func (h *GetSecondHandByTimeLine) loadData() error {
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

func (h *GetSecondHandByTimeLine) transToClientMessage() []*second_hand.SecondHandInfo {
	messageInfos := make([]*second_hand.SecondHandInfo, 0, len(h.messages))
	for _, messageItem := range h.messages {
		var uris []string
		if err := json.Unmarshal([]byte(messageItem.ImageUris), &uris); err != nil {
			logs.Error("message_id = %v image_uris = %v", messageItem.MessageId, messageItem.ImageUris)
			continue
		}
		userInfo, find := h.userInfo[messageItem.OpenId]
		if !find {
			logs.Warn("[GetSecondHandByTimeLine] transToClientMessage openId = %v not find UserInfo", messageItem.OpenId)
			continue
		}
		messageInfo := &second_hand.SecondHandInfo{
			MessageId:  messageItem.MessageId,
			Content:    messageItem.Content,
			Urls:       uris,
			CreateTime: messageItem.CreateTime.Unix(),
			Price:      messageItem.Money,
			Category:   messageItem.Category,
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
		messageInfo.DiggCount = int32(h.mid2DiggCount[messageItem.MessageId])
		messageInfos = append(messageInfos, messageInfo)
	}
	return messageInfos
}

func (h *GetSecondHandByTimeLine) GetOpenIDs(messages []model.SecondHand) []string {
	opendIds := make([]string, 0, len(messages))
	for _, mess := range messages {
		opendIds = append(opendIds, mess.OpenId)
	}
	return opendIds
}

func (h *GetSecondHandByTimeLine) loadUserInfo() error {
	openIDs := h.GetOpenIDs(h.messages)
	userInfo, err := db.MGetUserInfo(openIDs)
	if err != nil {
		return err
	}
	h.userInfo = userInfo
	return nil
}

func (h *GetSecondHandByTimeLine) loadMessages() error {
	var (
		messages []model.SecondHand
		err      error
	)
	if h.r.Category == "" {
		messages, err = db.GetSecondHandTimeLine(h.r.FirstTime, h.r.Index, h.r.Count+1)
	} else {
		messages, err = db.GetSecondHandTimeLineByCategory(h.r.FirstTime, h.r.Index, h.r.Count+1, h.r.Category)
	}
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

func (h *GetSecondHandByTimeLine) getMessageIds(messages []model.SecondHand) []int64 {
	messageIds := make([]int64, 0, len(messages))
	for _, mess := range messages {
		messageIds = append(messageIds, mess.MessageId)
	}
	return messageIds
}

func (h *GetSecondHandByTimeLine) loadDiggCount() error {
	messageIDs := h.getMessageIds(h.messages)
	diggCount, err := db.MGetDiggCount(messageIDs)
	if err != nil {
		return err
	}
	h.mid2DiggCount = diggCount
	return nil
}
