// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/DemoForService/controllers"
	_ "github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/updateUMsg", &controllers.UserController{}, "post:UpdateUMsg"),
		beego.NSRouter("/getUMsg/:account", &controllers.UserController{}, "get:GetUMsg"),
		beego.NSRouter("/getRecord/:account", &controllers.PayController{}, "get:GetRecord"),
		beego.NSRouter("/getAverage/:account", &controllers.PayController{}, "get:GetAverage"),
		beego.NSRouter("/downloadAvatar/:account", &controllers.UserController{}, "get:DownloadAvatar"),
		beego.NSRouter("/recharge", &controllers.UserController{}, "post:Recharge"),
		beego.NSRouter("/getAll/:all", &controllers.UserController{}, "get:GetAll"),


		beego.NSRouter("/upgrade", &controllers.UpgradeController{}, "post:Upload"),
		beego.NSRouter("/compare", &controllers.CompareController{}, "post:Compare"),
		beego.NSRouter("/getComparePic/:uid/:type", &controllers.PayController{}, "get:GetComparePic"),
		beego.NSRouter("/getCMsg/:uid", &controllers.CompareMsgController{}, "get:GetCMsg"),
		beego.NSRouter("/findPasswd", &controllers.UserController{}, "post:FindPasswd"),
		beego.NSRouter("/user_ctl/verify/:email", &controllers.VerifyController{}, ),
		beego.NSRouter("/searchRecord", &controllers.SearchController{},"post:SearchRecord"),
		beego.NSRouter("/searchUsers", &controllers.UserController{},"post:SearchUsers"),
		beego.NSRouter("/deleteUser", &controllers.UserController{},"post:DeleteUser"),
		beego.NSRouter("/updatePassword", &controllers.UserController{},"post:UpdatePassword"),
		beego.NSRouter("/create", &controllers.UserController{},"get,post:Create"),
	)
	beego.AddNamespace(ns)
}
