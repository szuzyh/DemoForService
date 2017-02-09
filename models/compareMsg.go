package models

import (
	_ "time"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "strconv"
)

func InsertIntoCompareMsg(account ,confidence,uid string, cMsg map[string]string){
	o := orm.NewOrm()
	sql := "insert into comparemsg(account,confidence,name,sex,birthday,nation,ID,address,start_date,end_date,department,pay_id,img_path) values('" + account + "','"+ confidence + "','"+ cMsg["personName"] + "','"+ cMsg["sex"] + "','"+ cMsg["birthday"] + "','"+ cMsg["nation"] + "','"+ cMsg["personId"] + "','"+ cMsg["address"] + "','"+ cMsg["startDate"] + "','"+ cMsg["endDate"] + "','"+ cMsg["department"] + "','"+ uid +"','"+ cMsg["imgPath"] + "')"
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
func QueryCMsgWithUid(uid string)(cmsgJson CMsg){
	o := orm.NewOrm()
	var sql string
	sql = "select  *from comparemsg where pay_id='"+uid+"'"
	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)
	cmsgJson.PersonName=p[0][3].(string)
	cmsgJson.Sex=p[0][4].(string)
	cmsgJson.Birthday=p[0][5].(string)
	cmsgJson.Nation=p[0][6].(string)
	cmsgJson.PersonId=p[0][7].(string)
	cmsgJson.Address=p[0][8].(string)
	cmsgJson.StartDate=p[0][9].(string)
	cmsgJson.EndDate=p[0][10].(string)
	cmsgJson.Department=p[0][11].(string)
	cmsgJson.ImgPath=p[0][13].(string)
	return
}
type CMsg struct {
	PersonName string `json:"personName"`
	Sex string `json:"sex"`
	Nation string `json:"nation"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	PersonId string `json:"personId"`
	Department string `json:"department"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	ImgPath string `json:"imgpath"`

}
