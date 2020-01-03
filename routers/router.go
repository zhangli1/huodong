package routers

import (
	"huodong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.MainController{}, "get:Add")
	beego.Router("/search", &controllers.MainController{}, "get:Search")
	beego.Router("/getLastMessage", &controllers.MainController{}, "get:GetLastMessage")
	beego.Router("/addMessage", &controllers.MainController{}, "post:AddMessage")
}
