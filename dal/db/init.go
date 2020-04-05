package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"my_project/conf"
	"my_project/logs"
)

var (
	db *gorm.DB
)
var localKey = "root:bing1120@tcp(127.0.0.1:3306)/app?charset=utf8"

func getRemoteKey() string {
	mysql := conf.Conf.GetMysql()
	return fmt.Sprintf("%s:%s@%s(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", mysql.UserName, mysql.Password, mysql.Network, mysql.Ip, mysql.Port, mysql.DbName)
}
func init() {
	var err error
	remoteKey := localKey
	logs.Info(remoteKey)
	db, err = gorm.Open("mysql", remoteKey)
	if err != nil {
		panic(err)
	}
}
