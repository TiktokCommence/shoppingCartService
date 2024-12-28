package data

import (
	"ShoppingCartService/internal/conf"
	"testing"
)

var config = &conf.Mysql{
	Database: &conf.Mysql_Database{
		Dsn: "root:123456@tcp(127.0.0.1:3307)/carts?charset=utf8mb4&parseTime=True&loc=Local",
	},
}

func TestNewMysql(t *testing.T) {
	if _, got := NewMysql(config); got != nil {
		t.Errorf("NewMysql() = %v, want %v", got, nil)
	}
}
