package controllers

import (
	"crypto/md5"
	"favgo/models"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"strconv"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login/login.tpl"
}

func (this *LoginController) Logout() {
	globalSessions.SessionDestroy(this.Ctx.ResponseWriter, this.Ctx.Request)

	this.Ctx.Redirect(302, "/")
}

//登陆处理
func (this *LoginController) Post() {

	_ = this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	user,err := models.UserAccountGetByDomain(username)

	this.Data["json"] = map[string]interface{}{"status": 0, "info": "参数错误"}


	if err ==nil {
		// 取MD5值
		md5Password := md5.New()
		io.WriteString(md5Password, password+ strconv.FormatInt(user.Uid,16))
		newPass := fmt.Sprintf("%x", md5Password.Sum(nil))


		// 密码验证
		if user.Passwordaes==newPass {
			sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
			defer sess.SessionRelease(this.Ctx.ResponseWriter)

			sess.Set("uid", user.Uid)
			sess.Set("userInfo", user)

			this.Data["json"] = map[string]interface{}{"status": 1, "info": "OK"}
		}else{
			this.Data["json"] = map[string]interface{}{"status": 0, "info": "密码不正确！"}
		}

		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(username,newPass,now)
		fmt.Println("user:=====")
		fmt.Println(user)
	}

	//this.Ctx.Redirect(302, "/")

	this.ServeJSON()

}
