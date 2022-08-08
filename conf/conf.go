package conf

import (
	"time"

	"github.com/go-ini/ini"
	logging "github.com/sirupsen/logrus"
)

var (
	HttpPort string

	MysqlHost            string
	MysqlPort            string
	MysqlUser            string
	MysqlPassword        string
	MysqlName            string
	MysqlIsLog           bool
	MysqlIsSingularTable bool
	MysqlMaxIdleConns    int
	MysqlMaxOpenConns    int
	MysqlConnMaxLifetime time.Duration
)

func Init(path string) {
	file, err := ini.Load(path)
	if err != nil {
		logging.Info("Fail to parse 'conf/app.ini': ", err)
	}

	loadService(file)
	loadMysql(file)
}

func loadService(file *ini.File) {
	HttpPort = file.Section("service").Key("HttpPort").MustString(":8080")
}

func loadMysql(file *ini.File) {
	section, err := file.GetSection("mysql")
	if err != nil {
		logging.Info(err)
	}
	MysqlHost = section.Key("MysqlHost").String()
	MysqlPort = section.Key("MysqlPort").String()
	MysqlUser = section.Key("MysqlUser").String()
	MysqlPassword = section.Key("MysqlPassword").String()
	MysqlName = section.Key("MysqlName").String()
	MysqlIsLog = section.Key("MysqlIsLog").MustBool(true)
	MysqlIsSingularTable = section.Key("MysqlIsSingularTable").MustBool(true)
	MysqlMaxIdleConns = section.Key("MysqlMaxIdleConns").MustInt(20)
	MysqlMaxOpenConns = section.Key("MysqlMaxOpenConns").MustInt(100)
	MysqlConnMaxLifetime = time.Duration(section.Key("MysqlConnMaxLifetime").MustInt(30)) * time.Second
}
