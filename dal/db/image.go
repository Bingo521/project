package db

import (
	"my_project/model"
)

func SaveImage(uri string,imgType string)error{
	image := model.Image{}
	image.Uri = uri
	image.Type = imgType
	return db.Save(&image).Error
}

func GetImageInfo(uri string)(*model.Image,error){
	image := model.Image{}
	err:=db.Where("uri = ?",uri).First(&image).Error
	if err != nil{
		return nil, err
	}
	return &image,nil
}