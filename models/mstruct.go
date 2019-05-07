package models

type Item struct {
	Id int `:"id"`
	Uid int64 `:"uid"`
	Title string `:"title"`
	Sort int `:"sort"`
	Description string `:"description"`
	IsTemplate int `:"is_template"`
	Page int `:"page"`
	Link_list []Link `orm:"-"`
}

type Link struct {
	Id int `:"id"`
	ItemId int `:"item_id"`
	Name string `:"name"`
	Href string `:"href"`
	Description string `:"description"`
	Seq int `:"seq"`
}
type Pwd_item struct {
	Uid int64 `:"uid"`
	Id int64 `:"id"`
	Title string `:"title"`
	Url string `:"url"`
	Description string `:"description"`
	Stat int `:"stat"`
	Slt string `:"slt"`
	Pwd string `:"pwd"`
	IsEncode int `:"is_encode"`
	Cid int `:"cid"`
	PromptMsg string `:"prompt_msg"`
	DisplayPwd int `:"display_pwd"`
	Seq int `:"seq"`
	Inserttime string `:"inserttime"`
	Updatetime string `:"updatetime"`
}
type User_account struct {
	Uid int64 `orm:"column(uid);pk"`
	AccessType int `:"access_type"` // 接入類型，1=>新浪
	Passwordaes string `:"passwordaes"`
	Level int `:"level"`
	Nickname string `:"nickname"`
	Domain string `:"domain"`
	Location string `:"location"`
	Province int `:"province"` // 省
	City int `:"city"`
	Gender string `:"gender"` // 性别
	Description string `:"description"`
	OauthToken string `:"oauth_token"`
	OauthTokenSecret string `:"oauth_token_secret"`
	AvatarsPath string `:"avatars_path"` // 頭像path
	UserEmail string `:"user_email"`
	AccessToken string `:"access_token"`
}
