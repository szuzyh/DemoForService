package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	_"github.com/DemoForService/docs"
	_ "github.com/DemoForService/routers"
)

func init() {
	err := orm.RegisterDataBase("default", "mysql", "root:passwd@tcp(127.0.0.1:3306)/facial?charset=utf8")
	if err != nil {
		panic(err)
	}
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
