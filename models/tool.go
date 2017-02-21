package models

import (
	"encoding/base64"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	_ "time"
	_ "fmt"

	"fmt"
	"time"
)

/*
* aColumns []string `SQL Columns to display`
* thismodel interface{} `SQL model to use`
* ctx *context.Context `Beego ctx which contains httpcontext`
* maps []orm.Params `return result in a interface map as []orm.Params`
* count int64 `return iTotalDisplayRecords`
* counts int64 `return iTotalRecords`
*
 */

type Datatable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func Datatables(aColumns []string, thismodel interface{}, Input *context.BeegoInput) (maps []orm.Params, count int64, counts int64) {
	/*
	 * Paging
	 * 分页请求
	 iDisplayStart  起始数目
	 iDisplayLength 每页显示数量
	*/
	iDisplayStart, _ := strconv.Atoi(Input.Query("iDisplayStart"))
	iDisplayLength, _ := strconv.Atoi(Input.Query("iDisplayLength"))


	/*
	 * Ordering
	 * 排序请求
	 */
	/*querysOrder := []string{}
	//if iSortCol_0 := tables["iSortCol_0"]; iSortCol_0 > -1 {
	if iSortCol_0 := 1; iSortCol_0 > -1 {
		ranges, _ := strconv.Atoi(Input.Query("iSortingCols"))
		for i := 0; i < ranges; i++ {
			istring := strconv.Itoa(i)
			if iSortcol := Input.Query("bSortable_" + Input.Query("iSortCol_"+istring)); iSortcol == "true" {
				sordir := Input.Query("sSortDir_" + istring)
				thisSortCol, _ := strconv.Atoi(Input.Query("iSortCol_" + istring))
				if sordir == "asc" {
					querysOrder = append(querysOrder, aColumns[thisSortCol])
				} else {
					querysOrder = append(querysOrder, "-"+aColumns[thisSortCol])
				}
			}
		}
	}*/
	/*
		       * Filtering
			        * 快速过滤器
	*/
	//querysFilter := []string{}
	cond := orm.NewCondition()
	if len(Input.Query("sSearch")) > 0 {
		for i := 0; i < len(aColumns); i++ {
			cond = cond.Or(aColumns[i]+"__icontains", Input.Query("sSearch"))

		}
	}

	if len(Input.Query("account")) > 0 {
		cond = cond.Or("account__icontains", Input.Query("account"))
	}


	/*timestamp begin to end*/
	//end, _ := strconv.Atoi(Input.Query("end"))
	//begin, _ := strconv.Atoi(Input.Query("begin"))
	//if begin > end {
	//	logs.Warn("Error: Invalid time. Begin must smaller than end")
	//} else {
	//	if begin != 0 || end != 0 {
	//		begins := time.Unix(int64(begin), 0)
	//		cond = cond.And("created__gte", begins.Format("2006-01-02 03:04:05"))
	//		ends := time.Unix(int64(end), 0)
	//		cond = cond.And("created__lte", ends.Format("2006-01-02 03:04:05"))
	//	}
	//
	//}

	/* Individual column filtering */
	for i := 0; i < len(aColumns); i++ {
		if Input.Query("bSearchable_"+strconv.Itoa(i)) == "true" && len(Input.Query("sSearch_"+strconv.Itoa(i))) > 0 {
			cond = cond.And(aColumns[i]+"__icontains", Input.Query("sSearch"))
		}
	}
	fmt.Println("")
	if len(Input.Query("min")) >0 {
		min,_ :=strconv.Atoi(Input.Query("min"))
		if Input.Query("min")=="NaN"{
			min=0
		}
		fmt.Println(min)
		cond = cond.And("confidence__gte", min)
	}
	if len(Input.Query("max")) > 0{
		max,_ :=strconv.Atoi(Input.Query("max"))
		if Input.Query("max")=="NaN"{
			max=100
		}
		fmt.Println(max)
		cond = cond.And("confidence__lte", max)
	}

	if len(Input.Query("start_time")) > 0{
		start_time :=Input.Query("start_time")
		t := ChangeNumToTime(start_time)
		fmt.Println(t)
		cond = cond.And("created__gte", t)
	}

	if len(Input.Query("end_time")) > 0{
		end_time :=Input.Query("end_time")
		var t string
		if end_time=="NaN"{
			end_time=strconv.Itoa(int(time.Now().Unix()))
			t = ChangeNumToTime(end_time)
		}else {
			tF := ChangeNumToTime(end_time)
			t1,_ := time.Parse("2006-01-02 15:04:05", tF)
			t1 =t1.Add(24*time.Hour)
			t = t1.Format("2006-01-02 15:04:05")
		}

		fmt.Println(t)
		cond = cond.And("created__lte", t)
	}
	//choose ipc
	//ipc := Input.Query("ipc")
	//if ipc != "" {
	//	ipc, _ := strconv.Atoi(ipc)
	//	cond = cond.And("ipc", ipc)
	//}

	/*
	 * GetData
	 * 数据请求
	 */
	o := orm.NewOrm()
	qs := o.QueryTable(thismodel)
	counts, _ = qs.Count()
	qs = qs.OrderBy("-"+aColumns[0])
	qs = qs.Limit(iDisplayLength, iDisplayStart)
	qs = qs.SetCond(cond)


	//	for _, v := range querysOrder {
	//	qs = qs.OrderBy(v)
	//}
	qs.Values(&maps)
	count, _ = qs.Count()
	return maps, count, counts
}

func SexMsgDatas(name string) ([]byte, error) {
	f, err := os.Open("/root/go/src/statistics/views/image/male.jpg")
	if err != nil {
		logs.Warn(err)
		return nil, err
	}
	defer f.Close()
	//生成base64
	sourcebuffer := make([]byte, 500000)
	n, _ := f.Read(sourcebuffer)
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	sexBase64 := "data:image/jpg;base64," + sourcestring

	return []byte(sexBase64), nil
}
func ChangeNumToTime(str string)(tFinal string){
	i,_ :=strconv.Atoi(str)
	fmt.Println(i)
	t := time.Unix(int64(i),0)
	tFinal =t.Format("2006-01-02 15:04:05")
	return
}
