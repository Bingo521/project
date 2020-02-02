package conf

import (
	"my_project/model"
	"my_project/util"
)

var Conf *model.Config
var confPath = "conf/conf.yml"
func init(){
	var err error
	Conf,err=util.ReadYamlFile(confPath)
	if err != nil{
		panic("load conf err")
	}
}