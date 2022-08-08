package main

import (
	"gin-gorm2.0/conf"
	"gin-gorm2.0/db/mysql"
	"gin-gorm2.0/route"
	logging "github.com/sirupsen/logrus"
)

func main() {
	r := route.NewRoute()
	err := r.Run(conf.HttpPort)
	if err != nil {
		logging.Fatalln(err)
	}
}

func init() {
	conf.Init("./conf/config.ini")
	mysql.Init()
}
