package conf

import (
	"encoding/json"
	"my_project/logs"
	"my_project/model"
	"my_project/util"
)

var Conf *model.Config
var confPath = "conf/conf.yml"

func init() {
	var err error
	Conf, err = util.ReadYamlFile(confPath)
	if err != nil {
		panic(err)
	}
	b,err:=json.Marshal(Conf)
	if err != nil{
		logs.Info("conf = %v",string(b))
	}
}
