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
	c *gin.Context
	comm *model.CtxComm
	req *message.GetUserMessageRequest
}

func NewGetUserMessageHandler(c *gin.Context)*GetUserMessageHandler{
	comm := util.GetCtxComm(c)
	return &GetUserMessageHandler{
		c: c,
		comm: comm,
	}
}

func (h * GetUserMessageHandler)Execute()*message.GetUserMessageResonse{
	resp := message.GetUserMessageResonse{}
	resp.StatusCode = 0
	resp.Message = "success"
	req,err := h.getReq()
	if err != nil {
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	h.req = req
	if h.req == nil {
		return &resp
	}
	if h.comm == nil{
		logs.Error("params is illegal")
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	messages,err := db.GetMessageByOpenId(h.req.OpenId,h.req.Index,h.req.Count+1)
	if err != nil{
		logs.Error("get message err:%v",err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = error_code.SYS_MESSAGE_SERVER_ERR
		return &resp
	}
	hasMore := false
	if len(messages) > int(h.req.Count){
		hasMore = true
		messages = messages[:h.req.Count]
	}
	messInfos := make([]*message.MessageInfo,0,len(messages))
	for i,_:=range messages{
		mess := &messages[i]
		messInfo := new(message.MessageInfo)
		messInfo.Content = mess.Message
		messInfo.MessageId = mess.MessageId
		var uris []string
		err:=json.Unmarshal([]byte(mess.ImageUris),&uris)
		if err != nil{
			logs.Error("Unmarshal ImageUris = %v err=%v",mess.ImageUris,uris)
			continue
		}
		for i,uri := range uris{
			uris[i] = "https://cmyk-so.com/"+imageOrigin(uri)
		}
		messInfo.Urls = uris
		messInfo.CreateTime = mess.CreateTime.Unix()
		messInfos = append(messInfos, messInfo)
	}
	resp.MessageInfos = messInfos
	resp.HasMore = hasMore
	return &resp
}


func (h *GetUserMessageHandler)getReq()(*message.GetUserMessageRequest,error){
	if h.c == nil{
		return nil,errors.New("context is nil")
	}
	strIndex := h.c.Query("index")
	strCount := h.c.Query("count")
	openId := h.c.Query("open_id")
	logs.Info("index = %v,count = %v,open_id = %v",strIndex,strCount,openId)
	if strIndex == "" || strCount == "" || openId == ""{
		return nil,fmt.Errorf("param not find")
	}
	index ,err := strconv.ParseInt(strIndex,10,64)
	if err != nil{
		return nil,fmt.Errorf("index is err:%v",err)
	}
	count,err := strconv.ParseInt(strCount,10,64)
	if err != nil{
		return nil,fmt.Errorf("count is err:%v" ,err)
	}
	if count == 0{
		return nil,nil
	}
	return &message.GetUserMessageRequest{
		Index: index,
		Count: count,
		OpenId: openId,
	},nil
}