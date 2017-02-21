package controllers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/models"
	"fmt"
	_ "time"
	"strconv"
)

type SearchController struct {
	beego.Controller
}


func (this *SearchController)SearchRecord(){
	fmt.Printf("%+v", string(this.Ctx.Input.RequestBody))
	payModel :=[]string{
		"Created",
		"Level",
		"Message",
		"Id",
	}
	maps, count, counts := models.Datatables(payModel, new(models.Pay), this.Ctx.Input)

	data := make(map[string]interface{}, count)
	var output = make([][]interface{}, len(maps))
	for i, m := range maps {
		for _, v := range payModel {
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
