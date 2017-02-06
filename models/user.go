package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"math/rand"
	"strings"
)


type Register struct {
	Id        int       `orm:"column(id);auto"`
	User      string    `orm:"column(user);size(64);null"`
	Password  string    `orm:"column(password);size(64);null"`
	Realname  string    `orm:"column(realname);size(200);null"`
	Tel       string    `orm:"column(tel);size(64);null"`
	Email     string    `orm:"column(email);size(64);null"`
	Hotelname string    `orm:"column(hotelname);size(64);null"`
	Location  string    `orm:"column(location);size(64);null"`
	Userlevel string    `orm:"column(userlevel);size(10);null"`
	Remain    string    `orm:"column(remain);size(10);null"`
	Created   string `orm:"column(created);type(datetime);null"`
	Total     string    `orm:"column(total);size(20)"`
	Avatar    string    `orm:"column(avatar);null"`
	Account   string    `orm:"column(account);size(10)"`
	Sex       string    `orm:"column(sex);size(10)"`
	Average   string    `orm:"column(average);size(10)"`
	Recharge  string    `orm:"column(recharge);size(10)"`
}

func (t *Register) TableName() string {
	return "register"
}
func init() {
	orm.RegisterModel(new(Register))
}

func AddRegister(user, password, email string) (account string, err error) {
	o := orm.NewOrm()
	t := time.Now().Format("2006-01-02 15:04:05")
	t = string(t)
	account = createAccount();
	if len(account) != 6 {
		account = account + "0"
	}
	r, err := o.Raw("insert into register(user,password,realname,tel,email,hotelname,location,userlevel,remain,created,total,account,sex,average,recharge) values('" + user + "','" + password + "','','','" + email + "','','','1','100','" + t + "','0','" + account + "','','0','0')").Exec()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		num, _ := r.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return
	}

}
func createAccount() (account string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	account = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return
}

func QueryIsExist(account string)(exist bool){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "select password from register where email='"+account+"'"
	}else {
		sql = "select password from register where account='"+account+"'"
	}
	var p []orm.Params
	o.Raw(sql).Values(&p)
	if p==nil{
		return false
	}
	return true
}
func QueryPassword(account string)(password string){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "select password from register where email='"+account+"'"
	}else {
		sql = "select password from register where account='"+account+"'"
	}

	var p []orm.Params
	o.Raw(sql).Values(&p)
	if str,ok := p[0]["password"].(string); ok {
		/* act on str */
		password = str
		return
	} else {
		/* not string */
		return "0"
	}
}
func QueryUMsg(account string)(user Register){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "select  *from register where email='"+account+"'"
	}else {
		sql = "select  *from register where account='"+account+"'"
	}

	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)
	user=ListToRegisterWithIndex(user,p,0)
	return
}

func UpdateUMsg(account,user,realname,tel,hotelname,location,sex,email string)(){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "update register SET user="+"'"+user+"', realname="+"'"+realname+"', tel="+"'"+tel+"', email="+"'"+email+"', hotelname="+"'"+hotelname+"', sex="+"'"+sex+"', location="+"'"+location+"'"+" where email='"+account+"'"
	}else {
		sql = "update register SET user="+"'"+user+"', realname="+"'"+realname+"', tel="+"'"+tel+"', email="+"'"+email+"', hotelname="+"'"+hotelname+"', sex="+"'"+sex+"', location="+"'"+location+"'"+" where account='"+account+"'"
	}
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
func UpdateAverage(account string,average string)(){
	o := orm.NewOrm()
	var sql string
	sql="update register set average='"+average+"' where account='"+account+"'"
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
func UpdateRecharge(account,recharge string)(){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql="update  register set recharge='"+recharge+"' where email='"+account+"'"
	}else {
		sql="update  register set recharge='"+recharge+"' where account='"+account+"'"
	}
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
func QueryRecharge(account string) (recharge string) {
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "select recharge from register where email='"+account+"'"
	}else {
		sql = "select recharge from register where account='"+account+"'"
	}

	var p []orm.Params
	o.Raw(sql).Values(&p)
	if str,ok := p[0]["password"].(string); ok {
		/* act on str */
		recharge = str
		return
	} else {
		/* not string */
		return "0"
	}
}
func QueryAllUser()(userList []Register){
	o := orm.NewOrm()
	var sql string
	sql = "select  *from register "
	var p []orm.ParamsList
	o.Raw(sql).ValuesList(&p)
	for i:=0;i<len(p);i++{
		var user Register
		user=ListToRegisterWithIndex(user,p,i)
		userList=append(userList,user)
	}

	return
}
func ListToRegisterWithIndex(user Register,p []orm.ParamsList,i int)(u Register){
	user.User = p[i][1].(string)
	user.Realname = p[i][3].(string)
	user.Tel = p[i][4].(string)
	user.Email= p[i][5].(string)
	user.Hotelname = p[i][6].(string)
	user.Location = p[i][7].(string)
	user.Userlevel = p[i][8].(string)
	user.Remain = p[i][9].(string)
	user.Created = p[i][10].(string)
	user.Total = p[i][11].(string)
	user.Account = p[i][13].(string)
	user.Sex = p[i][14].(string)
	user.Average = p[i][15].(string)
	user.Recharge = p[i][16].(string)
	return user
}
func QueryRemain(account string)(remain string){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql = "select remain from register where email='"+account+"'"
	}else {
		sql = "select remain from register where account='"+account+"'"
	}

	var p []orm.Params
	o.Raw(sql).Values(&p)
	if str,ok := p[0]["remain"].(string); ok {
		/* act on str */
		remain = str
		return
	} else {
		/* not string */
		return "0"
	}
}
func UpdateRemain(account,remain string){
	o := orm.NewOrm()
	var sql string
	endswith := strings.HasSuffix(account,".com")
	if endswith {
		sql="update  register set remain='"+remain+"' where email='"+account+"'"
	}else {
		sql="update  register set remain='"+remain+"' where account='"+account+"'"
	}
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