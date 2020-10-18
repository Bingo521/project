package handler

import (
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

type DeleteSecondHandHandler struct {
	c    *gin.Context
	req  *second_hand.DeleteSecondHandRequest
	comm *model.CtxComm

	messageInfo *model.Message
}

func NewDeleteSecondHandHandler(c *gin.Context) *DeleteSecondHandHandler {
	return &DeleteSecondHandHandler{
		c: c,
	}
}

func (h *DeleteSecondHandHandler) Execute() *second_hand.DeleteSecondHandResponse {
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

func (h *DeleteSecondHandHandler) makeReq() error {
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
	h.req = &second_hand.DeleteSecondHandRequest{
		MessageId: messageID,
	}
	return nil
}

func (h *DeleteSecondHandHandler) check() error {
	if h.req == nil || h.req.MessageId <= 0 {
		return fmt.Errorf("req err")
	}
	return nil
}

func (h *DeleteSecondHandHandler) loadData() error {
	messageInfo, err := db.GetMessageByMessageId(h.req.MessageId)
	if err != nil {
		return err
	}
	h.messageInfo = messageInfo
	return nil
}

func (h *DeleteSecondHandHandler) checkAfterLoad() error {
	if h.messageInfo == nil {
		return fmt.Errorf("message_id = %v not find", h.req.MessageId)
	}
	return nil
}

func (h *DeleteSecondHandHandler) delete() error {
	return db.DeleteSecondHandByMessageId(h.req.MessageId)
}

func (h *DeleteSecondHandHandler) makeErrResp(errNo int32, errMessage string) *second_hand.DeleteSecondHandResponse {
	return &second_hand.DeleteSecondHandResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}
