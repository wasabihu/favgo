package main

import (
	//"favgo/models"
	_ "favgo/routers"
	//"favgo/controllers"
	"github.com/astaxie/beego"
	//"fmt"

	 _ "favgo/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)


	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 10
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:123456@/www.wasa.ren?charset=utf8", maxIdle, maxConn)
}

func main() {

	beego.AddFuncMap("molding",molding)
	beego.AddFuncMap("modend",modend)

	beego.SetStaticPath("/images","public/images")
	beego.SetStaticPath("/styles","public/styles")
	beego.SetStaticPath("/js","public/js")
	beego.SetStaticPath("/static","public/static")
	beego.SetStaticPath("/lib","public/lib")

	beego.SetStaticPath("/indexbb","public/indexbb")
	beego.SetStaticPath("/logo","public/logo")



	beego.Run()


}

func molding(a int) int{
	return a%4
}

func modend(a int) int{
	if a !=0 && a+1%4==0{
		return 1
	}
	return 0
}


