package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"my_project/conf"
)
var (
	db *gorm.DB
)
var localKey = "root:cocos123456.@tcp(127.0.0.1:3306)/test_db?charset=utf8"

func getRemoteKey()string{
	mysql:=conf.Conf.GetMysql()
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8",mysql.UserName,mysql.Password,mysql.Network,mysql.Ip,mysql.Port,mysql.DbName)
}
func init() {
	var err error
	remoteKey := getRemoteKey()
	db, err = gorm.Open("mysql", remoteKey)
	if err != nil {
		panic("connect db err")
	}
}
