package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/comment"
	"my_project/util"
	"strconv"
)

type DeleteCommentHandler struct {
	c    *gin.Context
	req  *comment.DeleteCommentRequest
	comm *model.CtxComm

	messageInfo *model.Comment
}

func NewDeleteCommentHandler(c *gin.Context) *DeleteCommentHandler {
	return &DeleteCommentHandler{
		c: c,
	}
}

func (h *DeleteCommentHandler) Execute() *comment.DeleteCommentResponse {
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

func (h *DeleteCommentHandler) makeReq() error {
	comm := util.GetCtxComm(h.c)
	if comm == nil {
		return errors.New("comm info nil")
	}
	h.comm = comm
	strMessageID := h.c.PostForm("comment_id")
	logs.Info("[DeleteMessageHandler] comment_id = %v", strMessageID)
	messageID, err := strconv.ParseInt(strMessageID, 10, 64)
	if err != nil {
		return err
	}
	h.req = &comment.DeleteCommentRequest{
		CommentId: messageID,
	}
	return nil
}

func (h *DeleteCommentHandler) check() error {
	if h.req == nil || h.req.CommentId <= 0 {
		return fmt.Errorf("req err")
	}
	return nil
}

func (h *DeleteCommentHandler) loadData() error {
	messageInfo, err := db.GetCommentByCommentId(h.req.CommentId)
	if err != nil {
		return err
	}
	h.messageInfo = messageInfo
	return nil
}

func (h *DeleteCommentHandler) checkAfterLoad() error {
	if h.messageInfo == nil {
		return fmt.Errorf("message_id = %v not find", h.req.CommentId)
	}
	return nil
}

func (h *DeleteCommentHandler) delete() error {
	return db.DeleteSecondHandByMessageId(h.req.CommentId)
}

func (h *DeleteCommentHandler) makeErrResp(errNo int32, errMessage string) *comment.DeleteCommentResponse {
	return &comment.DeleteCommentResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}
