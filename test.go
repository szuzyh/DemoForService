package main



import (
	_ "fmt"

	_ "time"

	_ "strings"
	"fmt"
	"time"
)



func main() {
	//s:="123456_441424199508272258_head.jpg"
	//str := strings.Split(s,"_")
	//fmt.Println(str[0])
	//fmt.Println(str[1])
	//fmt.Println(str[2])
	//if  strings.Contains(s,"head"){
	//	fmt.Println("sss")
	//}

	//s:="2017-02-05 15:04:10"
	//tm2, _ := time.Parse("2006/01/02 03:04:05", s)
	//
	//fmt.Println(tm2.Format("2006/01/02 03"))
	////获取时间戳
	//
	//timestamp := time.Now().Unix()
	//
	//fmt.Println(timestamp)
	//
	//
	//
	////格式化为字符串,tm为Time类型
	//
	//tm := time.Unix(timestamp, 0)
	//
	//fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	//
	//fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
	//
	//
	//
	//
	//
	////从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	//
	//tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	//
	//
	////1486179075
	//tm3,_ := time.Parse("2006-01-02 03:04:05 PM","2017-02-04 11:31:15 AM")
	//fmt.Println(tm2.Unix())
	//fmt.Println(tm3)
	//
	toBeCharge := "2017-02-05 15:04:10"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	timeModel:= "2006-01-02 15"
	dataTimeStr := time.Unix(sr, 0).Format(timeModel) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}
