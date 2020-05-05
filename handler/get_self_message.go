package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/message"
	"my_project/util"
	"strconv"
)

type GetSelfMessageHandler struct {
	c *gin.Context
	comm *model.CtxComm
	req *message.GetSelfMessageRequest
}

func NewGetSelfMessageHandler(c *gin.Context)*GetSelfMessageHandler{
	comm := util.GetCtxComm(c)
	return &GetSelfMessageHandler{
		c: c,
		comm: comm,
	}
}

func (h * GetSelfMessageHandler)Execute()*message.GetSelfMessageResonse{
	h.req = h.getReq()
	resp := message.GetSelfMessageResonse{}
	resp.StatusCode = 0
	resp.Message = "success"
	if h.req == nil || h.comm == nil{
		logs.Error("params is illegal")
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = error_code.SYS_MESSAGE_PARAM_ILLEGAL
		return &resp
	}
	messages,err := db.GetMessageByOpenId(h.comm.OpenId,h.req.Index,h.req.Count+1)
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


func (h *GetSelfMessageHandler)getReq()*message.GetSelfMessageRequest{
	if h.c == nil{
		return nil
	}
	strIndex := h.c.Query("index")
	strCount := h.c.Query("count")
	logs.Info("index = %v,count = %v",strIndex,strCount)
	if strIndex == "" || strCount == ""{
		return nil
	}
	index ,err := strconv.ParseInt(strIndex,10,64)
	if err != nil{
		return nil
	}
	count,err := strconv.ParseInt(strCount,10,64)
	if err != nil{
		return nil
	}
	if count == 0{
		return nil
	}
	return &message.GetSelfMessageRequest{
		Index: index,
		Count: count,
	}
}