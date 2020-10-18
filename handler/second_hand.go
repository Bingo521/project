package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/second_hand"
	"my_project/util"
	"strconv"
)

type SecondHandHandler struct {
	c    *gin.Context
	comm *model.CtxComm
	req  *second_hand.CreateSecondHandRequest
}

func NewSecondHandHandler(c *gin.Context) *SecondHandHandler {
	comm := util.GetCtxComm(c)
	return &SecondHandHandler{
		c:    c,
		comm: comm,
	}
}

func (h *SecondHandHandler) Execute() *second_hand.CreateSecondHandResponse {
	resp := second_hand.CreateSecondHandResponse{}
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
		logs.Error("[SecondHandHandler] get req err:%v", err)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	if err := h.checkParams(); err != nil {
		logs.Error("[SecondHandHandler] checkParams err:%v", err)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	userInfo, err := db.GetUserInfo(openId)
	if err != nil {
		logs.Warn("[MessageHandler] openId = %v getUserInfo err:%v", openId, err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	mId, err := redis.Inc("message_id")
	if err != nil {
		logs.Error("[SecondHandHandler] get message ID err:%v", err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	mess, err := db.CreateSecondHand(openId, mId, h.req.Content, h.req.Uris, h.req.Price, h.req.Category)
	if err != nil {
		logs.Error("[SecondHandHandler] create message err :%v", err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}

	resp.MessageInfo = &second_hand.SecondHandInfo{
		Content:      h.req.Content,
		Urls:         h.req.Uris,
		Category:     h.req.Category,
		Price:        h.req.Price,
		MessageId:    mId,
		CreateTime:   mess.CreateTime.Unix(),
		DiggCount:    0,
		CommentCount: 0,
		Digg:         false,
		UserInfo:     transToSecondHandInfo(userInfo),
	}
	return &resp
}

func (h *SecondHandHandler) getReq() (*second_hand.CreateSecondHandRequest, error) {
	content := h.c.PostForm("content")
	strUris := h.c.PostForm("uris")
	price := h.c.PostForm("price")
	category := h.c.PostForm("category")
	logs.Info("[MessageHandler] content = %v,uris = %v,price = %v,category = %v", content, strUris, price, category)

	var uris []string
	err := json.Unmarshal([]byte(strUris), &uris)
	if err != nil {
		logs.Error("get req err:uris = %v,err = %v", strUris, err)
		return nil, err
	}
	fPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		logs.Warn("[MessageHandler] price = %v is illegal", price)
		return nil, err
	}
	return &second_hand.CreateSecondHandRequest{
		Content:  content,
		Uris:     uris,
		Price:    float32(fPrice),
		Category: category,
	}, nil
}

func (h *SecondHandHandler) checkParams() error {
	if h.req == nil {
		return errors.New("param is nil")
	}
	if h.req.Content == "" || len(h.req.Uris) == 0 || h.req.Category == "" {
		return errors.New("param illegal")
	}
	return nil
}

func transToSecondHandInfo(userInfo *model.UserInfo) *second_hand.SecondHandUserInfo {
	if userInfo == nil {
		return nil
	}
	localUserInfo := second_hand.SecondHandUserInfo{}
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
