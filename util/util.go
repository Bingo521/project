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

func Retry(f func()error,t int)error{
	if t <= 0{
		return nil
	}
	var err error
	for t > 0{
		t--
		err=f()
		if err == nil{
			return nil
		}
	}
	return err
}