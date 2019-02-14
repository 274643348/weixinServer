package routers

import (
	"weixin/goServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/down1", "static")
	beego.SetStaticPath("/down2", "static2")
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
}
