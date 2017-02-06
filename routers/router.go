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
		beego.NSRouter("/ifaces", &controllers.UserController{}, "get:Ifaces"),
		beego.NSRouter("/getAll/:all", &controllers.UserController{}, "get:GetAll"),
		beego.NSNamespace("/compare", beego.NSInclude(&controllers.CompareController{}, ), ),
		beego.NSNamespace("/upgrade", beego.NSInclude(&controllers.UpgradeController{}, ), ),
		beego.NSRouter("/getComparePic/:uid/:type", &controllers.PayController{}, "get:GetComparePic"),

	)
	beego.AddNamespace(ns)
}
