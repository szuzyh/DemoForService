package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type EmailVerify struct {
	Email string `json:"email"`
	Created string `json:"created"`
	User string `json:"user"`
	Password string `json:"password"`
	Location string `json:"location"`
}

func QueryAllWithEmail(email string)(verify EmailVerify){
	o := orm.NewOrm()
	fmt.Println("email:"+email)
	sql := "select *from emailverify where email = '"+email+"'"
	var p []orm.Params
	o.Raw(sql).Values(&p)
	fmt.Println(p)
	if p==nil{
		verify.Email=""
		verify.Created=""
		verify.User=""
		verify.Password=""
		verify.Location=""
		return
	}else {
		verify.Email=p[0]["email"].(string)
		verify.Created=p[0]["created"].(string)
		verify.User=p[0]["user"].(string)
		verify.Password=p[0]["password"].(string)
		verify.Location=p[0]["location"].(string)
		return
	}
}
func QueryCreated(email string)(created string){
	o := orm.NewOrm()
	sql := "select created from emailverify where email = '"+email+"'"
	var p []orm.Params
	o.Raw(sql).Values(&p)
	if p==nil{
		created=""
		return
	}else {
		created=p[0]["created"].(string)
		return
	}
}
func DeleteEmail(email string)(error error){
	o := orm.NewOrm()
	sql := "delete  from emailverify where email = '"+email+"'"
	_,err := o.Raw(sql).Exec()
	if err !=nil{
		return err
	}
	return nil
}
func InsertVerify(email,user,password,location string)(){
	o := orm.NewOrm()
	t := time.Now().Format("2006-01-02 15:04:05")
	t = string(t)


	r, err := o.Raw("insert into emailverify(email,created,user,password,location) values('" + email + "','" + t + "','" + user + "','"+password+"','" + location+ "')").Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		num, _ := r.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return
	}

}
