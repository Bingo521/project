package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/comment"
)

type CommentHandler struct {
	c   *gin.Context
	req *comment.CreateCommentRequest
}

func NewCommentHandler(c *gin.Context) *CommentHandler {
	return &CommentHandler{
		c: c,
	}
}

func (h *CommentHandler) Execute() *comment.CreateCommentResponse {
	req, err := h.makeReq()
	if err != nil {
		logs.Warn("[CommentHandler] makeReq err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_PARAM_ILLEGAL)
	}
	h.req = req
	mId, err := redis.Inc("message_id")
	if err != nil {
		logs.Warn("[CommentHandler] gen id err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	if err := db.CreateComment(mId, req.CommentInfo.MessageId, req.CommentInfo.Content, req.CommentInfo.Uris, model.COMMENT_TYPE_COMMENT); err != nil {
		logs.Warn("[CommentHandler] create err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	return h.makeResp()
}

func (h *CommentHandler) makeReq() (*comment.CreateCommentRequest, error) {
	commentInfo := h.c.PostForm("comment_info")
	logs.Info("[CommentHandler] comment_info = %v", commentInfo)
	commentMainInfo := comment.CommentMainInfo{}
	if err := json.Unmarshal([]byte(commentInfo), &commentMainInfo); err != nil {
		logs.Warn("[CommentHandler] makeReq Unmarshal err:%v", err)
		return nil, err
	}
	req := comment.CreateCommentRequest{}
	req.CommentInfo = &commentMainInfo
	return &req, nil
}

func (h *CommentHandler) makeErrResp(errNo int32, errMessage string) *comment.CreateCommentResponse {
	return &comment.CreateCommentResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}

func (h *CommentHandler) makeResp() *comment.CreateCommentResponse {
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.CommentInfo = h.req.CommentInfo
	return resp
}
