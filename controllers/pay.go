package controllers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/models"

	_ "fmt"
	"strconv"
)

type PayController struct {
	beego.Controller
}

func (c *PayController)GetRecord(){
	defer c.ServeJSON()
	account := c.GetString(":account")
	var arr JsonM
	if !models.QueryIsExist(account){
		arr.Detail= "account not found"
		arr.Status="failed"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}
	var list []models.Pay
	list = models.QueryRecord(account)
	var j JsonM2
	for i:=0;i<len(list);i++{
		var pay models.Pay
		pay=list[i]
		var j3 JsonM3
		j3.Account=pay.Account
		j3.Message=pay.Message
		j3.Created=pay.Created
		j3.Level=pay.Level
		j.Detail=append(j.Detail,j3)
	}
	j.Status="success"
	c.Data[`json`]= j
	c.ServeJSON()
	return
}
type JsonM2 struct {
	Detail []JsonM3 `json:"detail"`
	Status string `json:"status"`
}
type JsonM3 struct {
	Account string `json:"account"`
	Created string  `json:"created"`
	Message string  `json:"message"`
	Level string  `json:"level"`
}

type Message struct {
	num string `json:"num"`
}
func setRecordData(Map map[string]string, pay models.Pay) {
	Map["account"]=pay.Account
	Map["created"]=pay.Created
	Map["message"]=pay.Message
	Map["level"]=pay.Level
}

func(c *PayController)GetAverage(){
	defer c.ServeJSON()
	account := c.GetString(":account")
	var arr JsonM
	average := models.QueryMessage(account)
	models.UpdateAverage(account,strconv.Itoa(average))
	arr.Detail= strconv.Itoa(average)
	arr.Status="success"
	c.Data[`json`]= arr
	c.ServeJSON()
	return
}
