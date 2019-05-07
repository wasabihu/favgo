package controllers

import (
	"favgo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"net/http/pprof"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	"sync"
)

var itemLsMap sync.Map
var globalSessions *session.Manager
//var uid int64 = 1659969325
//var uid int64 = 0


type MainController struct {
	beego.Controller
}

type ProfController struct {
	beego.Controller
}


func GetSessUid(ctx *context.Context) int64{
	var uid int64 = 0
	if uid <= 0 {
		sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
		defer sess.SessionRelease(ctx.ResponseWriter)

		var v = sess.Get("uid")
		if v != nil {
			uid = v.(int64)
		}
	}
	fmt.Println("session uid:",uid)

	return uid
}

func (c *MainController) Get() {

	var lists =  []models.Item{}

	uid := GetSessUid(c.Ctx)

	if uid >0{
		sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		defer sess.SessionRelease(c.Ctx.ResponseWriter)

		c.Data["is_edit"] = true

		v := sess.Get("userInfo")
		if v != nil {
			userInfo := v.(models.User_account)
			c.Data["user_info"] = userInfo
		}

	} else{
		c.Data["is_edit"] = false
		c.Data["user_info"] = models.User_account{}

		if c.GetString("u")=="wasa"{
			uid = 1659969325
		}

	}

	vv, ok := itemLsMap.Load(uid)

	if ok {
		lists,ok = vv.([]models.Item)
	}

	if !ok{
		var mutex sync.Mutex

		mutex.Lock()

		var err error
		lists,_,err = models.ItemLsGet(uid,200)

		if err == nil {
			for k, row := range lists {
				lists[k].Link_list, _, _ = models.GetLinks(row.Id)
			}
		}else {
			fmt.Println(err)
		}

		fmt.Println("Store:",uid)
		itemLsMap.Store(uid,lists)

		mutex.Unlock()
	}

	c.Data["user_id"] =uid

	c.Data["item_length"] =len(lists)
	c.Data["item_list"] = lists

	c.Data["curr_page"] = "/"
	c.TplName = "index.tpl"
}




func (this *ProfController) Get() {

	var pp string = this.Ctx.Input.Param(":pp")

	switch pp{
		default:
			pprof.Index(this.Ctx.ResponseWriter, this.Ctx.Request)
		case "":
			pprof.Index(this.Ctx.ResponseWriter, this.Ctx.Request)
		case "cmdline":
			pprof.Cmdline(this.Ctx.ResponseWriter, this.Ctx.Request)
		case "profile":
			pprof.Profile(this.Ctx.ResponseWriter, this.Ctx.Request)
		case "symbol":
			pprof.Symbol(this.Ctx.ResponseWriter, this.Ctx.Request)
	}

	//pprof.Index(this.Ctx.ResponseWriter, this.Ctx.Request)
	//pprof.Profile(this.Ctx.ResponseWriter, this.Ctx.Request)
	this.Ctx.ResponseWriter.WriteHeader(200)
}


func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}