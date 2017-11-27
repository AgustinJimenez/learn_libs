package xorm_mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//_ = engine.Sync2(new(Users))
func GetInstance() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@/iris?charset=utf8")
	if err != nil {
		fmt.Println("ERROR ON xorm_mysql GetInstance()", err)
	}

	return engine
}
