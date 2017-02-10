package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/DemoForService/models"
	"time"
)

type VerifyController struct {
	beego.Controller
}
func (c *VerifyController)Get(){

	email :=c.GetString(":email")
	created :=models.QueryCreated(email)
	t :=time.Now()
	createdT,err := time.Parse("2006-01-02 15:04:05", created)
	if err!=nil{
		fmt.Println(err.Error())
	}else {
		createdAdd := createdT.Add(24*time.Hour)
		if t.Before(createdAdd){
			verify := models.QueryAllWithEmail(email)
			if verify.Location=="" {
				c.Ctx.Output.Body([]byte("邮箱错误"))
				return
			}
			account, err := models.AddRegister(verify.User, verify.Password, email,verify.Location)
			if err != nil {
				c.Ctx.Output.Body([]byte("错误:"+err.Error()))
				return
			} else {
				err :=models.DeleteEmail(email)
				if err!=nil {
					c.Ctx.Output.Body([]byte("错误:"+err.Error()))
					return
				}
				c.Ctx.Output.Body([]byte("成功验证,你的账户是"+account))
				return
			}

		}else {
			c.Ctx.Output.Body([]byte("链接已过期"))
			return
		}
	}


	return
}

//func (c * VerifyController)Get(){
//	c.TplName="verify.html"
//}