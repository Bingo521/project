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
)

type MessageHandler struct {
	c *gin.Context
	comm *model.CtxComm
	req *message.CreateMessageRequest
}

func NewMessageHandler(c *gin.Context)*MessageHandler{
	comm := util.GetCtxComm(c)
	return &MessageHandler{
		c: c,
		comm: comm,
	}
}

func (h *MessageHandler)Execute()*message.CreateMessageResponse{
	resp := message.CreateMessageResponse{}
	resp.StatusCode = 0
	resp.Message = "success"
	if h.c == nil || h.comm == nil{
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	var err error
	openId := h.comm.OpenId
	h.req ,err = h.getReq()
	if err != nil{
		logs.Error("get req err:%v",err)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	mId,err:=redis.Inc("message_id")
	if err != nil{
		logs.Error("get message ID err:%v",err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	mess,err:=db.CreateMessage(openId,mId,consts.MESSAGE_TYPE_THREAD,h.req.Content,h.req.Uris)
	if err != nil{
		logs.Error("create message err :%v",err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	resp.Content = h.req.Content
	resp.Uris = h.req.Uris
	resp.MessageId = mId
	resp.CreateTime = mess.CreateTime.Unix()
	return &resp
}

func (h *MessageHandler)getReq()(*message.CreateMessageRequest,error){
	content := h.c.PostForm("content")
	strUris := h.c.PostForm("uris")
	var uris []string
	err := json.Unmarshal([]byte(strUris),&uris)
	if err != nil{
		logs.Error("get req err:uris = %v,err = %v",strUris,err)
		return nil,err
	}
	return &message.CreateMessageRequest{
		Content: content,
		Uris: uris,
	},nil
}

func (h *MessageHandler)checkParams()error{
	if h.req == nil{
		return errors.New("param is nil")
	}
	if h.req.Content == "" || len(h.req.Uris) == 0{
		return errors.New("param illegal")
	}
	return nil
}
