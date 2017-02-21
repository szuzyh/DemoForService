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

	_ "html/template"

	"time"
)

type UserController struct {
	beego.Controller
	ds DummySession
}
type DummySession struct {
	Cnt int
}
//注册
func (c *UserController) Register() {
	defer c.ServeJSON()
	user := c.GetString("user")
	password := c.GetString("password")
	email := c.GetString("email")
	location := c.GetString("location")
	arr := make(map[string]string)
	if models.QueryIsEmailExist(email){
		arr["detail"] = "email is used"
		arr["status"] = "failed"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}
	//models.InsertVerify(email,user,password,location)
	//models.SendVerifyMail(email,"测试链接http://120.76.128.35:4569/api/user_ctl/verify/"+email)
	account, err := models.AddRegister(user, password, email,location)
	if err!=nil {
		arr["detail"] = err.Error()
		arr["status"] = "failed"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}else {
		arr["detail"] = account
		arr["status"] = "success"
		c.Data[`json`]= arr
		c.ServeJSON()
		return
	}
	//arr["detail"] = "send email,please verify"
	//arr["status"] = "success"
	//controller.Data[`json`]= arr
	//controller.ServeJSON()
	//return


}
//登录

func (c *UserController)Login(){
	defer c.ServeJSON()
	password := c.GetString("password")
	account := c.GetString("account")

	var arr JsonM
	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
	if account=="0" {
		arr.Detail="account not found"
		arr.Status="failed"
	}else {
		if !models.QueryIsExist(account) {
			arr.Detail = "account not found"
			arr.Status = "failed"
		} else {
			//if strings.EqualFold(models.QueryOnline(account), "true") {
			//	arr.Detail = "account online"
			//	arr.Status = "failed"
			//} else {
				str := models.QueryPassword(account)
				s := models.CreateSHA(str,"/login")
				fmt.Println(s)
				if !strings.EqualFold(password, s) {
					arr.Detail = "password error"
					arr.Status = "failed"
				} else {
					token := models.CreateMD5(account+time.Now().Format("2006-01-02 15:04:05"))
					arr.Detail = token
					arr.Status = "success"
					models.UpdateOnline(account, "true")
					if is,_ :=models.QueryLoginToken(account);!is {
						models.AddToken(account,token)
					}else {
						models.UpdateLoginToken(account,token)
					}

					c.SetSession("userid",account)


					//arr.Detail="login"

				}
			}

		//}
	}
	c.Data[`json`]= arr
	c.ServeJSON()
	return
}
func (c *UserController)GetUMsg(){
	defer c.ServeJSON()
	account := c.GetString(":account")

	//m :=c.Ctx.Request.Header.Get("Cookie")
	//m1 := models.StrToMap(m)
	//fmt.Println(m1["userid"])
	//fmt.Println(m1["beegosessionID"])
	fmt.Println(c.GetSession("userid"))
	//fmt.Println(c.GetSession("beegosessionID"))

	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
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
	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
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
	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
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
	jall.Status="success"
	jall.Detail=userARecord
	c.Data[`json`]= jall
	c.ServeJSON()
	return
}

type JsonAll struct {
	Detail UserAndRecord `json:"detail"`
	Status string `json:"status"`
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
	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
	ff, err := os.Open("/tmp/account/"+account+"/head/"+account+"_head.jpg")
	if err!=nil{
		fmt.Println(err)
		ff,_ =os.Open("/tmp/account/null.jpg")
	}
	defer ff.Close()
	//生成base64
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	s := "data:image/jpg;base64,"+sourcestring
	c.Ctx.Output.Body([]byte(s))
}
func (c *UserController)FindPasswd(){
	defer c.ServeJSON()
	email :=c.GetString("email")
	//token :=c.GetString("token")
	//if err:=models.CheckToken(token);err!=nil{
	//	json :=JsonM{Detail:"token error"+err.Error(),Status:"failed"}
	//	c.Data[`json`]=json
	//	c.ServeJSON()
	//	return
	//}
	//fmt.Println(email)
	if !(strings.Contains(email,".")||strings.Contains(email,"@")){
		email = models.QueryEmailWithAccount(email)
	}
	if email=="0"{
		json :=JsonM{Detail:"email error",Status:"failed"}
		c.Data[`json`]=json
		c.ServeJSON()
		return
	}
	if !models.QueryIsEmailExist(email){
		json :=JsonM{Detail:"email error",Status:"failed"}
		c.Data[`json`]=json
		c.ServeJSON()
		return
	}
	password :=models.QueryPassword(email)
	err :=models.SendMail(email,password)
	if err!=""{
		json :=JsonM{Detail:err,Status:"failed"}
		c.Data[`json`]=json
		c.ServeJSON()
		return
	}
	json :=JsonM{Detail:"password send to email",Status:"success"}
	c.Data[`json`]=json
	c.ServeJSON()
	return
}

func (this *UserController)SearchUsers(){
	fmt.Printf("%+v", string(this.Ctx.Input.RequestBody))
	UserModel :=[]string{
		"Created",
		"Account",
		"User",
		"Location",
		"Id",
	}
	maps, count, counts := models.Datatables(UserModel, new(models.Register), this.Ctx.Input)

	data := make(map[string]interface{}, count)
	var output = make([][]interface{}, len(maps))
	for i, m := range maps {
		for _, v := range UserModel {
			output[i] = append(output[i], m[v])

			//if v == "Uid" {
			//	output[i] = append(output[i],m[v].())
			//}else {
			//	output[i] = append(output[i], m[v])
			//}
		}
	}

	data["sEcho"], _ = strconv.Atoi(this.Ctx.Input.Query("sEcho"))
	data["iTotalRecords"] = counts
	data["iTotalDisplayRecords"] = count
	data["aaData"] = output
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *UserController)DeleteUser(){

	account := this.GetString("account")
	callback :=models.DeleteUser(account)
	var j JsonM
	if callback!="1"{
		j.Detail=callback
		j.Status="failed"
		this.Data["json"]=j
		this.ServeJSON()
		return
	}
	j.Detail="delete!"
	j.Status="success"
	this.Data["json"]=j
	this.ServeJSON()
	return
}
func (this *UserController)UpdatePassword(){
	var j JsonM
	account := this.GetString("account")
	old := this.GetString("old")
	new := this.GetString("new")
	fmt.Println(old)
	if strings.Contains(account,".")||strings.Contains(account,"@"){
		account = models.QueryAccountWithEmail(account)
	}
	if account=="0"{
		j.Detail="account error"
		j.Status="failed"
		this.Data["json"]=j
		this.ServeJSON()
		return
	}

	if old!=models.QueryPassword(account){
		j.Detail="old password error"
		j.Status="failed"
		this.Data["json"]=j
		this.ServeJSON()
		return
	}
	callback :=models.UpdatePass(account,new)

	if callback!="1"{
		j.Detail=callback
		j.Status="failed"
		this.Data["json"]=j
		this.ServeJSON()
		return
	}
	j.Detail="update!"
	j.Status="success"
	this.Data["json"]=j
	this.ServeJSON()
	return
}
const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)
type RoomData struct {
	Messages        []Message
	RemainingPeople int
}
var (
	_     = fmt.Printf // To prevent compiler complaining about unused imports
	data  = make(map[string]*RoomData)
	rooms = make([]string, 20) // Contains name of every currently active room
)
func contains(rooms []string, roomName string) bool {
	for _, name := range rooms {
		if name == roomName {
			return true
		}
	}
	return false
}
func redirectWithError(controller *UserController,
	errorMessage string,
	path string) {

	flash := beego.NewFlash()
	flash.Error(errorMessage)
	flash.Store(&controller.Controller)
	controller.Redirect(path, 302)
}
func (controller *UserController) Create() {
	controller.TplName = "create.html"
	beego.ReadFromRequest(&controller.Controller)
	if controller.Ctx.Input.Method() == METHOD_POST {
		roomName := controller.GetString("room-name")
		if contains(rooms, roomName) {
			redirectWithError(controller, "This room already exists.", "/create")
		} else {
			username := controller.GetString("username")
			temp := make(map[string]interface{})
			temp["roomName"] = roomName
			temp["username"] = username
			controller.SetSession(roomName, temp)
			rooms = append(rooms, roomName)
			data[roomName] = &RoomData{make([]Message, 0, 0), 1}
			controller.Redirect("/room/"+roomName, 302)
		}
	}
}