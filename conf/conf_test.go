package conf

import (
	"fmt"
	"testing"
)

func init() {
	Init("./config.ini")
}

func TestConf(t *testing.T) {
	fmt.Println(HttpPort)
	fmt.Println(MysqlHost)
	fmt.Println(MysqlPort)
}
