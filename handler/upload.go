package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/dal/redis"
	"my_project/error_code"
	"my_project/logs"
	"my_project/proto_gen/upload"
	"strings"
)

type ImageHandler struct {
	c *gin.Context
}

func NewImageHandler(c *gin.Context)*ImageHandler{
	return &ImageHandler{
		c: c,
	}
}

func (h * ImageHandler)Execute()*upload.UploadImageResponse{
	resp := upload.UploadImageResponse{}
	resp.Message = "success"
	resp.StatusCode = 0
	if h.c == nil{
		logs.Error("c is nil")
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = "param is illegal"
		return &resp
	}
	file,err:=h.c.FormFile("image")
	if err != nil{
		logs.Error("FormFile(\"image\") fail!err=%v",err)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = "param is illegal"
		return &resp
	}
	strs:=strings.Split(file.Filename,".")
	imgType := strings.ToLower(strs[len(strs)-1])
	imgTypeIsOk :=  checkImgType(imgType)
	if !imgTypeIsOk{
		logs.Error("image type = %v is illegal",imgType)
		resp.StatusCode = error_code.ERR_PARAM_ILLEGAL
		resp.Message = "param is illegal"
		return &resp
	}
	fileName ,err := redis.Inc("image_uri")
	uri := fmt.Sprintf("%012x.%v",fileName,imgType)
	if err != nil{
		logs.Error("get uri fail,err = %v",err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = "service err"
		return &resp
	}
	path := imageOrigin(uri)
	err = h.c.SaveUploadedFile(file,"./"+path)
	if err != nil{
		logs.Error("save file err:err = %v",err)
		resp.StatusCode = error_code.ERR_SERVER_ERR
		resp.Message = "service err"
		return &resp
	}
	resp.Uri = uri
	resp.Url = "https://cmyk-so.com/"+path
	return &resp
}

func checkImgType(imgType string)bool{
	switch imgType {
	case "jpg","jpeg","png","gif":
		return true
	default:
		return false
	}
}

func imageOrigin(imageName string)string{
	return "source/image/origin/"+imageName
}