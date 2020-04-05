package model

type ErrResp struct {
	StatusCode           int32    `json:"status_code"`
	Message              string   `json:"message"`
}
