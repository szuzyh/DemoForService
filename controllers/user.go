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
	fmt.Println(account)
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
	fmt.Println(s)
	fmt.Println(s.User)
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
	fmt.Println("------------All Record --------")
	fmt.Println(recordList)

	fmt.Println("--------- All User --------")
	models.QueryAllUser()
	
}

type JsonAll struct {
	Detail map[string]UserAndRecord
	status string `json:"status"`
}
type JsonRecord struct {
	Account string `json:"account"`
	Created string  `json:"created"`
	Message string  `json:"message"`
	Level string  `json:"level"`
}
 type UserAndRecord struct {
	 Users []models.Register  `json:"users"`
	 Records []JsonRecord `json:"records"`
 }
