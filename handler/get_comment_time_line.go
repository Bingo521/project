package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my_project/dal/db"
	"my_project/error_code"
	"my_project/logs"
	"my_project/model"
	"my_project/proto_gen/comment"
	"strconv"
)

type GetCommentTimeLineHandler struct {
	c                   *gin.Context
	req                 *comment.GetCommentRequest
	commentID2DiggCount map[int64]int64
}

func NewGetCommentTimeLineHandler(c *gin.Context) *GetCommentTimeLineHandler {
	return &GetCommentTimeLineHandler{
		c: c,
	}
}

func (h *GetCommentTimeLineHandler) Execute() *comment.GetCommentResponse {
	req, err := h.makeReq()
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] makeReq err:%v", err)
		return h.makeErrResp(error_code.ERR_PARAM_ILLEGAL, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	commentInfo, err := db.GetCommentByTimeLine(req.MessageId, req.FirstTime, req.Index, req.Count+1)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] GetCommentByTimeLine err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	hasMore := false
	if len(commentInfo) > int(req.Count) {
		hasMore = true
		commentInfo = commentInfo[:req.Count]
	}
	clientCommentInfo, err := h.makeClientCommentInfo(commentInfo)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] makeClientCommentInfo err:%v", err)
		return h.makeErrResp(error_code.ERR_SERVER_ERR, error_code.SYS_MESSAGE_SERVER_ERR)
	}
	resp := h.makeErrResp(error_code.ERR_SUCCESS, error_code.SYS_MESSAGE_SUCCESS)
	resp.CommentInfos = clientCommentInfo
	resp.HasMore = hasMore
	resp.NextIndex = req.Index + int32(len(commentInfo))
	resp.FirstTime = req.FirstTime
	return resp
}

func (h *GetCommentTimeLineHandler) makeReq() (*comment.GetCommentRequest, error) {
	strMessageID := h.c.Query("message_id")
	strIndex := h.c.Query("index")
	strCount := h.c.Query("count")
	strFirstTime := h.c.Query("first_time")
	logs.Info("[GetCommentTimeLineHandler] message_id = %v,index = %v,count = %v,first_time = %v", strMessageID, strIndex, strCount, strFirstTime)
	messageID, err := strconv.ParseInt(strMessageID, 10, 64)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] message_id err:%v", strMessageID)
		return nil, err
	}
	index, err := strconv.ParseInt(strIndex, 10, 64)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] index err:%v", strIndex)
		return nil, err
	}
	count, err := strconv.ParseInt(strCount, 10, 64)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] count = %v", strCount)
		return nil, err
	}
	firstTime, err := strconv.ParseInt(strFirstTime, 10, 64)
	if err != nil {
		logs.Warn("[GetCommentTimeLineHandler] first_time = %v", strFirstTime)
		return nil, err
	}
	req := comment.GetCommentRequest{}
	req.Index = int32(index)
	req.Count = int32(count)
	req.MessageId = messageID
	req.FirstTime = firstTime
	return &req, nil
}

func (h *GetCommentTimeLineHandler) makeErrResp(errNo int32, errMessage string) *comment.GetCommentResponse {
	return &comment.GetCommentResponse{
		StatusCode: errNo,
		Message:    errMessage,
	}
}

func (h *GetCommentTimeLineHandler) makeClientCommentInfo(commentInfos []model.Comment) ([]*comment.CommentInfo, error) {
	clientCommentInfos := make([]*comment.CommentInfo, 0, len(commentInfos))
	for _, commetItem := range commentInfos {
		var uris []string
		if err := json.Unmarshal([]byte(commetItem.ImageUris), &uris); err != nil {
			logs.Warn("[GetCommentTimeLineHandler] err:%v", err)
			continue
		}
		clientCommentItem := &comment.CommentInfo{
			CommentId:  commetItem.CommentID,
			DiggCount:  int32(commetItem.DiggCount),
			CreateTime: commetItem.CreateTime.Unix(),
			CommentMainInfo: &comment.CommentMainInfo{
				MessageId: commetItem.CommentID,
				Content:   commetItem.Content,
				Uris:      uris,
			},
		}
		clientCommentInfos = append(clientCommentInfos, clientCommentItem)
	}
	return clientCommentInfos, nil
}
