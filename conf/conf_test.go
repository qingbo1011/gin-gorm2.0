package conf

import (
	"testing"
)

func init() {
	Init("./config.ini")
}

func TestConf(t *testing.T) {
	//fmt.Println(HttpPort)
	//fmt.Println(MysqlHost)
	//fmt.Println(MysqlPort)
	//fmt.Println(MysqlLogMode)
	//fmt.Println(MysqlDataBase)
	//fmt.Println(MysqlSingularTable)
}
