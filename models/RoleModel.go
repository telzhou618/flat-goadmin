package models

type Role struct {
	Id int `pk:"auto" form:"id"`
	RoleName string `orm:"size(32)"`
	Remark string `orm:"size(300)"`
	User []*User `orm:"reverse(many)"`
}
