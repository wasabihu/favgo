package routers

import (
	"favgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/item", &controllers.ItemController{})

	beego.Router("/item", &controllers.ItemController{})

	beego.Router("/item/editlink",&controllers.ItemController{},"*:EditLink")
	beego.Router("/item/deletelink",&controllers.ItemController{},"*:DelLink")

	beego.Router("/item/edititem",&controllers.ItemController{},"*:EditItem")



	beego.Router(`/debug/pprof`, &controllers.ProfController{})
	beego.Router(`/debug/pprof/:pp([\w]+)`, &controllers.ProfController{})


	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/login", &controllers.LoginController{})

	beego.Router("/login/logout", &controllers.LoginController{},"*:Logout")

	//admin.Run()



}

//func FilterGetUid (ctx *context.Context) {
//
//}
