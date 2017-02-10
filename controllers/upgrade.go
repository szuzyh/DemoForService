package controllers

import (
	"os"
	"path"
	"github.com/astaxie/beego"
	"strings"
	_ "time"
	"github.com/DemoForService/models"
	_ "bytes"
	_ "image/jpeg"
	_ "fmt"
)

// Operations about Upgrade
type UpgradeController struct {
	beego.Controller
}

// @Title uploadFile
// @Description upload files
// @Param	
// @Success 200 {int} confidence
// @Failure 403 body is empty
// @router / [post]
func (u *UpgradeController) Post() {
	f, h, err := u.GetFile("imageName")                  //获取上传的文件
	if err != nil {
		u.Data["json"] = Result{err.Error(), "fail"}
		u.ServeJSON()
		return
	}

	var fpath string
	var account,Type string
	if strings.Contains(h.Filename,"head"){
		account = spiltHeadImg(h.Filename)
		if strings.Contains(account,".")||strings.Contains(account,"@"){
			account = models.QueryAccountWithEmail(account)
		}


		fpath = path.Join("/tmp/account/"+account+"/head", h.Filename)
		err = os.MkdirAll("/tmp/account/"+account+"/head", 0700)
	}else {
		account,Type = spiltImageName(h.Filename)
		fpath = path.Join("/tmp/compare/"+account+"/base", account+"_"+Type)
		err = os.MkdirAll("/tmp/compare/"+account+"/base", 0700)
	}

	if err != nil {
		u.Data["json"] = Result{err.Error(), "fail"}
		u.ServeJSON()
		return
	}

	f.Close()
	err = u.SaveToFile("imageName", fpath)
	if err != nil {
		u.Data["json"] = Result{err.Error(), "fail"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = Result{"Success", "success"}
	u.ServeJSON()
	return
}
func spiltImageName(imageName string)(account,Type string){
	str := strings.Split(imageName,"_")
	account=str[0]
	Type =str[2]
	return
}
func spiltHeadImg(imageName string)(account string){
	str := strings.Split(imageName,"_")
	account=str[0]
	return
}

