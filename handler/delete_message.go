package handler

import (
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

type DeleteMessageHandler struct {
	c    *gin.Context
	req  *message.DeleteMessageRequest
	comm *model.CtxComm

	messageInfo *model.Message
}

func NewDeleteMessageHandler(c *gin.Context) *DeleteMessageHandler {
	return &DeleteMessageHandler{
		c: c,
	}
}

func (h *DeleteMessageHandler) Execute() *message.DeleteMessageResponse {
	if err := h.makeReq(); err != nil {
		logs.Warn("[DeleteMessageHandler] makeReq err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.check(); err != nil {
		logs.Warn("[DeleteMessageHandler] check err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.loadData(); err != nil {
		logs.Warn("[DeleteMessageHandler] loadData err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	if err := h.checkAfterLoad(); err != nil {
		logs.Warn("[DeleteMessageHandler] checkAfterLoad err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	if err := h.delete(); err != nil {
		logs.Warn("[DeleteMessageHandler] delete err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	return h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
}

func (h *DeleteMessageHandler) makeReq() error {
	comm := util.GetCtxComm(h.c)
	if comm == nil {
		return errors.New("comm info nil")
	}
	h.comm = comm
	strMessageID := h.c.PostForm("message_id")
	logs.Info("[DeleteMessageHandler] message_id = %v", strMessageID)
	messageID, err := strconv.ParseInt(strMessageID, 10, 64)
	if err != nil {
		return err
	}
	h.req = &message.DeleteMessageRequest{
		MessageId: messageID,
	}
	return nil
}

func (h *DeleteMessageHandler) check() error {
	if h.req == nil || h.req.MessageId <= 0 {
		return fmt.Errorf("req err")
	}
	return nil
}

func (h *DeleteMessageHandler) loadData() error {
	messageInfo, err := db.GetMessageByMessageId(h.req.MessageId)
	if err != nil {
		return err
	}
	h.messageInfo = messageInfo
	return nil
}

func (h *DeleteMessageHandler) checkAfterLoad() error {
	if h.messageInfo == nil {
		return fmt.Errorf("message_id = %v not find", h.req.MessageId)
	}
	return nil
}

func (h *DeleteMessageHandler) delete() error {
	return db.DeleteMessageByMessageId(h.req.MessageId)
}

func (h *DeleteMessageHandler) makeErrResp(errNo int32, errMessage string) *message.DeleteMessageResponse {
	return &message.DeleteMessageResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}
