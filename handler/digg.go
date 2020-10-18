package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/digg"
	"my_project/util"
	"strconv"
)

type DiggHandler struct {
	c    *gin.Context
	comm *model.CtxComm
	req  *digg.DiggRequest
}

func NewDiggHandler(c *gin.Context) *DiggHandler {
	return &DiggHandler{
		c: c,
	}
}

func (h *DiggHandler) Execute() *digg.DiggResponse {
	req, err := h.makeReq()
	if err != nil {
		logs.Warn("[DiggHandler] err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.req = req
	if err := h.check(); err != nil {
		logs.Warn("[DiggHandler] chekc err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	switch h.req.DiggType {
	case digg.DiggType_DIGG:
		if err := db.Digg(h.comm.OpenId, h.req.MessageId); err != nil {
			logs.Warn("[DiggHandler] digg err:%v", err)
			return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
		}
	case digg.DiggType_CANCAL_DIGG:
		if err := db.DiggCancel(h.comm.OpenId, h.req.MessageId); err != nil {
			logs.Warn("[DiggHandler] cancel digg err:%v", err)
			return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
		}
	}
	count, err := db.GetDiggCountByMessageId(h.req.MessageId)
	if err != nil {
		logs.Warn("[DiggHandler] digg count err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	if err := db.SetDiggCount(h.req.MessageId, count); err != nil {
		logs.Warn("[DiggHandler] update digg count:%v", err)
	}
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.DiggCount = int32(count)
	return resp
}

func (h *DiggHandler) makeReq() (*digg.DiggRequest, error) {
	comm := util.GetCtxComm(h.c)
	if comm == nil {
		return nil, errors.New("comm info nil")
	}
	h.comm = comm
	strMessageId := h.c.PostForm("message_id")
	strType := h.c.PostForm("digg_type")
	logs.Info("[DiggHandler] open_id = %v,message_id = %v,diggType = %v", comm.OpenId, strMessageId, strType)
	messageId, err := strconv.ParseInt(strMessageId, 10, 64)
	if err != nil {
		return nil, err
	}
	diggType, err := strconv.ParseInt(strType, 10, 64)
	if err != nil {
		return nil, err
	}
	req := digg.DiggRequest{}
	req.MessageId = messageId
	req.DiggType = digg.DiggType(diggType)
	return &req, nil
}

func (h *DiggHandler) check() error {
	if h.req == nil {
		return errors.New("nil req")
	}
	if h.req.DiggType != digg.DiggType_DIGG && h.req.DiggType != digg.DiggType_CANCAL_DIGG {
		return errors.New("digg type err")
	}
	return nil
}

func (h *DiggHandler) makeErrResp(errNo int32, errMessage string) *digg.DiggResponse {
	return &digg.DiggResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}
