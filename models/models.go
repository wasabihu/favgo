package models

import (
	"github.com/astaxie/beego/orm"
)


func ItemLsGet(uid int64 ,limit int) ([]Item,int64,error) {
	o := orm.NewOrm()
	//qs := o.QueryTable("item")
	//var lists []orm.ParamsList
	var lists []Item
	//num,err := qs.Filter("uid", uid).OrderBy("sort").Limit(limit).ValuesList(&lists)

	num, err := o.Raw("SELECT * from item where uid=? order by sort limit ?", uid,limit).QueryRows(&lists)

	return lists,num,err
}

func GetLinks(item_id int) ([]Link,int64,error){
	o := orm.NewOrm()
	var lists []Link
	num, err := o.Raw("SELECT * from link where item_id=? order by seq", item_id).QueryRows(&lists)
	if err == nil {
		//fmt.Println("user nums: ", num)
	}
	return lists,num,err
}

func LinkAdd(link *Link) (int64, error){
	o := orm.NewOrm()
	return o.Insert(link)
}
func LinkEdit(link *Link) (int64, error){
	o := orm.NewOrm()
	return o.Update(link)
}

func LinkDel(link *Link) (int64, error){
	o := orm.NewOrm()
	return o.Delete(link)
}

func ItemAdd(item *Item) (int64, error){
	o := orm.NewOrm()
	return o.Insert(item)
}
func ItemEdit(item *Item) (int64, error){
	o := orm.NewOrm()
	return o.Update(item,"Title","Sort","Page")
}

//func CheckUserItemAcl(uid int64,itemId int) {
//}


func UserAccountGetByDomain(domain string) (user User_account, err error){
	o := orm.NewOrm()
	user = User_account{Domain: domain}
	err = o.Read(&user,"Domain")
	return
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Item), new(Link), new(Pwd_item), new(User_account))
}
