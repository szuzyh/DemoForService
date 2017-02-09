package controllers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/models"
)

type CompareMsgController struct {
	beego.Controller
}

func (c *CompareMsgController)GetCMsg(){
	defer c.ServeJSON()
	uid := c.GetString(":uid")
	cmsgJson:=models.QueryCMsgWithUid(uid)

	c.Data[`json`]=cmsgJson
	c.ServeJSON()
	return

}

type msgJson struct {
	Detail models.CMsg `json:"detail"`
	Status string `json:"status"`
}
