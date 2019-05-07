package controllers

import (
	"favgo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
)
type ItemController struct {
	beego.Controller
}

func (this *ItemController) DelLink() {
	uid := GetSessUid(this.Ctx)
	defer this.ServeJSON()

	if uid <1 {
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "用户未登录！"}
		return
	}
	var link_id int

	this.Ctx.Input.Bind(&link_id, "link_id")

	link := &models.Link{Id:link_id}

	valid := validation.Validation{}

	valid.Required(link.Id, "link_id")
	valid.Min(link.Id,1,"link_id")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "参数错误！"}
		return
	}

	models.LinkDel(link)

	// 刷新缓存
	itemLsMap.Delete(uid)

	this.Data["json"] = map[string]interface{}{"status": 1, "info": "OK"}
	fmt.Println("OK....",link)
	return
}

func (this *ItemController) EditLink() {
	uid := GetSessUid(this.Ctx)
	defer this.ServeJSON()

	if uid <1 {
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "用户未登录！"}
		return
	}

	var name,href,description string
	var link_id,seq,item_id int

	this.Ctx.Input.Bind(&name, "name")
	this.Ctx.Input.Bind(&href, "href")
	this.Ctx.Input.Bind(&description, "description")
	this.Ctx.Input.Bind(&seq, "seq")
	this.Ctx.Input.Bind(&item_id, "item_id")
	this.Ctx.Input.Bind(&link_id, "link_id")


	link := &models.Link{Name:name,Href:href,Description:description,Seq:seq,ItemId:item_id}

	valid := validation.Validation{}
	valid.Required(link.Name, "name")
	valid.Required(link.Href, "href")
	valid.Required(link.ItemId, "item_id")
	valid.Min(link.ItemId,1,"item_id")
	valid.Required(link.Seq, "seq")
	valid.Min(link.Seq, 0,"seq")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "参数错误！"}
		return
	}

	if link_id > 0{
		link.Id = link_id;
		_, err :=models.LinkEdit(link)
		if err != nil {
			fmt.Println(err)
		}
	}else{
		_, err :=models.LinkAdd(link)
		if err != nil {
			fmt.Println(err)
		}
	}
	// 刷新缓存
	itemLsMap.Delete(uid)

	this.Data["json"] = map[string]interface{}{"status": 1, "info": "OK"}
	fmt.Println("OK....",link)
	return

}

func (this *ItemController) EditItem() {
	uid := GetSessUid(this.Ctx)
	defer this.ServeJSON()

	if uid <1 {
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "没有权限！"}
		return
	}

	//var title string
	//var item_id,sort,page int

	item := &models.Item{Uid:uid}

	this.Ctx.Input.Bind(&item.Title, "title")
	this.Ctx.Input.Bind(&item.Id, "item_id")

	if item.Id < 1 {
		item.Page =1
		item.Sort =1
	}else{
		this.Ctx.Input.Bind(&item.Sort, "sort")
		this.Ctx.Input.Bind(&item.Page, "page")
	}

	valid := validation.Validation{}
	valid.Required(item.Title, "title")
	valid.Min(item.Uid,1,"uid")
	valid.Min(item.Page,1,"page")
	valid.Min(item.Sort, 0,"sort")


	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.Data["json"] = map[string]interface{}{"status": 0, "info": "参数错误！"}
		return
	}

	fmt.Println("item:",item)
	if item.Id > 0 {
		num, err := models.ItemEdit(item)
		if err != nil {
			fmt.Println(num,err)
		}
	}else{
		num, err := models.ItemAdd(item)
		if err != nil {
			fmt.Println(num,err)
		}
	}

	// 刷新缓存
	itemLsMap.Delete(uid)
	this.Data["json"] = map[string]interface{}{"status": 1, "info": "OK"}
	return
}


func (c *ItemController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}