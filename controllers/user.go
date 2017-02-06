package controllers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/models"

	"fmt"
	_ "encoding/json"
	_ "os"
	_ "strings"

	"strings"
	"strconv"
	_ "io/ioutil"
	"encoding/base64"
	"os"


)

type UserController struct {
	beego.Controller
}

//注册
func (controller *UserController) Register() {
	defer controller.ServeJSON()
	user := controller.GetString("user")
	password := controller.GetString("password")
	email := controller.GetString("email")
	arr := make(map[string]string)
	if models.QueryIsExist(email){
		arr["detail"] = "email is used"
		arr["status"] = "failed"
		controller.Data[`json`]= arr
		controller.ServeJSON()
		return
	}
	account, err := models.AddRegister(user, password, email)
	if err != nil {
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	} else {
		arr["detail"] = account
		arr["status"] = "success"
		controller.Data[`json`]= arr
		controller.ServeJSON()
		return
	}

}
//登录

func (c *UserController)Login(){
	defer c.ServeJSON()
	password := c.GetString("password")
	account := c.GetString("account")
	arr := make(map[string]string)
	if !models.QueryIsExist(account){

		arr["detail"] = "account not found"
		arr["status"] = "failed"

	}else {
		str := models.QueryPassword(account)
		if !strings.EqualFold(password,str){
			arr["detail"] = "password error"
			arr["status"] = "failed"
		}else {
			arr["detail"] = "login"
			arr["status"] = "success"
		}
	}
	c.Data[`json`]= arr
	c.ServeJSON()
}
func (c *UserController)GetUMsg(){
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
	var s models.Register
	s = models.QueryUMsg(account)
	Map := make(map[string]string)
	setData(Map,s)
	j := JsonM1{Detail:Map,Status:"success"}
	c.Data[`json`]= j
	c.ServeJSON()
}

func(c *UserController)UpdateUMsg(){
	defer c.ServeJSON()
	account := c.GetString("account")
	user := c.GetString("user")
	realname := c.GetString("realname")
	tel := c.GetString("tel")
	email := c.GetString("email")
	hotelname := c.GetString("hotelname")
	location := c.GetString("location")
	sex := c.GetString("sex")
	var arr JsonM
	if !models.QueryIsExist(account){
		arr.Detail= "account not found"
		arr.Status="failed"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}
	models.UpdateUMsg(account,user,realname,tel,hotelname,location,sex,email)
	arr.Detail= "update!"
	arr.Status="success"
	c.Data[`json`]= arr
	c.ServeJSON()
	return

}


func setData(m map[string]string, register models.Register) {
	m["user"]= register.User
	m["realname"]= register.Realname
	m["tel"]= register.Tel
	m["account"]= register.Account
	m["recharge"]= register.Recharge
	m["remain"]= register.Remain
	m["hotelname"]= register.Hotelname
	m["location"]= register.Location
	m["average"]= register.Average
	m["userlevel"]= register.Userlevel
	m["sex"]= register.Sex
	m["created"]= register.Created
	m["total"]= register.Total
	m["email"]= register.Email
}

type JsonM struct {
	Detail string `json:"detail"`
	Status string  `json:"status"`
}
type JsonM1 struct {
	Detail map[string]string `json:"detail"`
	Status string `json:"status"`
}

func (c *UserController)Recharge(){
	defer c.ServeJSON()
	account := c.GetString("account")
	sum := c.GetString("sum")
	var arr JsonM
	if !models.QueryIsExist(account){
		arr.Detail= "account not found"
		arr.Status="failed"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}
	r :=models.QueryRecharge(account)
	i, _ := strconv.Atoi(r)
	n, _ :=strconv.Atoi(sum)
	i+=n
	models.UpdateRecharge(account,strconv.Itoa(i))
	models.AddRecordRecharge(account,sum,strconv.Itoa(i))
	arr.Detail= "recharge"
	arr.Status="success"
	c.Data[`json`]= arr
	c.ServeJSON()
	return
}
func (c *UserController)Ifaces(){
	defer c.ServeJSON()
	var arr JsonM
	arr.Detail= "get ifaces"
	arr.Status="success"
	c.Data[`json`]= arr
	c.ServeJSON()
	return
}

func (c *UserController)GetAll(){
	defer c.ServeJSON()
	all := c.GetString(":all")
	fmt.Println(all)
	//var All JsonAll
	var recordList []models.Pay
	recordList = models.QueryAllRecord()
	var jall JsonAll
	var userARecord UserAndRecord
	for i:=0;i<len(recordList);i++{
		var pay models.Pay
		pay=recordList[i]
		var j3 JsonRecord
		j3.Uid=strconv.Itoa(pay.Id)
		j3.Account=pay.Account
		j3.Message=pay.Message
		j3.Created=pay.Created
		j3.Level=pay.Level
		j3.ID=pay.ID
		userARecord.Records=append(userARecord.Records,j3)
	}



	var userList []models.Register
	userList = models.QueryAllUser()
	for i:=0;i<len(userList);i++{
		var pay models.Register
		pay=userList[i]
		var j3 JsonUser
		j3.Account=pay.Account
		j3.Average=pay.Average
		j3.Created=pay.Created
		j3.Email=pay.Email
		j3.Hotelname=pay.Hotelname
		j3.Location=pay.Location
		j3.Realname=pay.Realname
		j3.Recharge=pay.Recharge
		j3.Tel=pay.Tel
		j3.Userlevel=pay.Userlevel
		j3.Sex=pay.Sex
		j3.Total=pay.Total
		j3.Userlevel=pay.Userlevel
		j3.Remain=pay.Remain
		userARecord.Users=append(userARecord.Users,j3)
	}
	jall.status="success"
	jall.Detail=userARecord
	c.Data[`json`]= jall
	c.ServeJSON()
	return
}

type JsonAll struct {
	Detail UserAndRecord `json:"detail"`
	status string `json:"status"`
}
type JsonRecord struct {
	Uid string `json:"uid"`
	Account string `json:"account"`
	Created string  `json:"created"`
	Message string  `json:"message"`
	Level string  `json:"level"`
	ID string `json:"id"`
}
 type UserAndRecord struct {
	 Users []JsonUser  `json:"users"`
	 Records []JsonRecord `json:"records"`
 }
type JsonUser struct {
	User string `json:"user"`
	Realname string `json:"realname"`
	Tel string `json:"tel"`
	Email string `json:"email"`
	Hotelname string `json:"hotelname"`
	Location string `json:"location"`
	Userlevel string `json:"userlevel"`
	Remain string `json:"remain"`
	Created string `json:"created"`
	Total string `json:"total"`
	Account string `json:"account"`
	Sex string `json:"sex"`
	Average string `json:"average"`
	Recharge string `json:"recharge"`
}

func(c *UserController)DownloadAvatar(){
	defer c.ServeJSON()
	account := c.GetString(":account")

	ff, err := os.Open("/tmp/account/"+account+"/head/"+account+"_head.jpg")
	if err!=nil{
		fmt.Println(err)
		ff,_ =os.Open("/tmp/account/null.jpg")
	}
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	//c.Ctx.Output.Download("/tmp/account/"+account+"/head/"+account+"_head.jpg",account+"_head.jpg")
	//c.Data[`string`]="data:image/jpg;base64,"+strings.TrimSpace(sourcestring)
	//c.Ctx.ResponseWriter.("data:image/jpg;base64,"+strings.TrimSpace(sourcestring))


	//s := "data:image/jpg;base64,"+strings.TrimSpace(sourcestring)
	s := "data:image/jpg;base64,"+sourcestring
	//c.Ctx.ResponseWriter.Write([]byte(s))
	//c.Ctx.Output.Context.WriteString()
	//c.Ctx.WriteString(s)
	//var arr JsonNull
	//arr.Detail=s;
	//arr.Status="success"
	//c.Data[`json`]=arr
	c.Ctx.Output.Body([]byte(s))
	//c.Data[`json`]=s
	//c.ServeJSON()
	//return
}