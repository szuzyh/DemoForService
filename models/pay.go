package models

import (
	"github.com/astaxie/beego/orm"
	_ "strings"
	_ "fmt"
	"fmt"
	"strings"
	"strconv"
	"time"
)

type Pay struct {
	Id        int       `orm:"column(uid);auto"`
	Account string      `orm:"column(account);size(10)"`
	Created string     `orm:"column(created);type(datetime);null"`
	Message string      `orm:"column(message);size(100)"`
	Level string      `orm:"column(level);size(20)"`
}

type PayList struct {
	Map map[string]string
}
func (t *Pay) TableName() string {
	return "pay"
}

func init() {
	orm.RegisterModel(new(Pay))
}

func QueryRecord(account string)(list []Pay){
	o := orm.NewOrm()
	var sql string
	sql = "select  *from pay where account='"+account+"'"
	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)

	for i:=0;i<len(p);i++{
		var pay Pay
		pay.Account=p[i][1].(string)
		pay.Created=p[i][2].(string)
		pay.Message=p[i][3].(string)
		pay.Level=p[i][4].(string)
		list=append(list,pay)
	}
	fmt.Println(list)
	return list
}
func QueryAllRecord()(list []Pay){
	o := orm.NewOrm()
	var sql string
	sql = "select  *from pay "
	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)
	for i:=0;i<len(p);i++{
		var pay Pay
		pay.Account=p[i][1].(string)
		pay.Created=p[i][2].(string)
		pay.Message=p[i][3].(string)
		pay.Level=p[i][4].(string)
		list=append(list,pay)
	}
	fmt.Println(list)
	return list
}
func setData(Map map[string]string, pay Pay) {
	Map["account"]=pay.Account
	Map["created"]=pay.Created
	Map["message"]=pay.Message
	Map["level"]=pay.Level
}
func QueryMessage(account string) (average int) {
	o := orm.NewOrm()
	var sql string
	sql = "select * from pay where account='"+account+"'and level='info'"
	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)
	var msg []string
	fmt.Println()
	if len(p)==0 {
		return 0
	}
	for i:=0;i<len(p);i++{
		if !strings.Contains(p[i][3].(string),"金额") {
			//msg[i]=p[i][3].(string)
			msg=append(msg,p[i][3].(string))
		}
	}
	for i:=0;i<len(msg);i++{
		str := strings.Split(msg[i],"相似度为:")
		s := strings.Split(str[1],",")
		a,err := strconv.Atoi(s[0])
		if err!=nil {
			fmt.Println("类型转化失败")
		}
		average+=a
	}

	return average/len(msg)
}
func AddRecordRecharge(account,recharge,sum string)(){
	o := orm.NewOrm()
	 msg:="用户:"+account+" 进行了一次充值,充值金额为:"+recharge+",账户金额为:"+sum;
	t := time.Now().Format("2006-01-02 15:04:05")
	t = string(t)
	sql := "insert into pay(account,created,message,level) values('" + account + "','" + t + "','" + msg + "','info')"
	r,err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		num, _ := r.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return
	}
}

