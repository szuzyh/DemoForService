package controllers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/models"
	"strconv"
	"fmt"
	"time"
	"strings"
	"os"
	"encoding/base64"
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
	if len(list)==0{
		var jsonNull JsonNull
		jsonNull.Detail=""
		jsonNull.Status="success"
		c.Data[`json`]= jsonNull
		c.ServeJSON()
		return
	}
	for i:=0;i<len(list);i++{
		var pay models.Pay
		pay=list[i]
		var j3 JsonM3
		j3.Uid=strconv.Itoa(pay.Id)
		j3.Account=pay.Account
		j3.Message=pay.Message
		j3.Created=pay.Created
		j3.Level=pay.Level
		j3.ID=pay.ID
		fmt.Println(j3.ID)
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
	Uid string `json:"uid"`
	Account string `json:"account"`
	Created string  `json:"created"`
	Message string  `json:"message"`
	Level string  `json:"level"`
	ID string `json:"id"`
}

type JsonNull struct {
	Detail string `json:"detail"`
	Status string `json:"status"`
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

func (c *PayController)GetComparePic() {
	defer c.ServeJSON()
	id := c.GetString(":uid")
	Type := c.GetString(":type")
	pay := models.QueryRecordWithUid(id)

	//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                              //转化所需模板
	loc, _ := time.LoadLocation("Local")                             //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, pay.Created, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                             //转化为时间戳 类型是int64
	                                                //打印输出时间戳 1420041600
	timeModel := "2006-01-02 15"
	dataTimeStr := time.Unix(sr, 0).Format(timeModel) //设置时间戳 使用模板格式化为日期字符串

	sourcebuffer := make([]byte, 500000)
	var sourcestring string
	if strings.EqualFold(Type, "person") {
		ff, err := os.Open("/tmp/compare/" + pay.Account + "/" + dataTimeStr + "/" + pay.Account + "_" + pay.ID + "_person.jpg")
		if err != nil {
			fmt.Println(err)
			ff, _ = os.Open("/tmp/compare/error.jpg")
		}
		defer ff.Close()
		n, _ := ff.Read(sourcebuffer)
		sourcestring = base64.StdEncoding.EncodeToString(sourcebuffer[:n])

	} else {
		ff, err := os.Open("/tmp/compare/" + pay.Account + "/" + dataTimeStr + "/" + pay.Account + "_" + pay.ID + "_own.jpg")
		if err != nil {
			fmt.Println(err)
			ff, _ = os.Open("/tmp/compare/error.jpg")
		}
		defer ff.Close()
		n, _ := ff.Read(sourcebuffer)
		sourcestring = base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	}
	s := "data:image/jpg;base64," + sourcestring
	c.Ctx.Output.Body([]byte(s))
}