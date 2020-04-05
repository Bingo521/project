package util

import (
	"gopkg.in/yaml.v2"
	"my_project/model"
	"os"
)

func ReadYamlFile(path string) (*model.Config, error) {
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}


func MakeErrResp(errCode int32,message string)*model.ErrResp{
	return &model.ErrResp{
		StatusCode:errCode,
		Message: message,
	}
}