package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int64	`pk:"auto"`
	UserName string `orm:"unique,size(30)" form:"userName" valid:"Required;MaxSize(20);MinSize(6)"`
	Password string `orm:"size(32)" form:"password"`
	Repassword string `orm:"-" form:"repassword"`
	Remark string `orm:"null;size(300)" form:"remark"`
	Status int `orm:"default(1)" form:"status"`
	CreateTime time.Time `orm:type(datetime)`
	Role []*Role `orm:"rel(m2m)"`
}

func  AddUser(user *User) (int64,error)  {
	user.CreateTime = time.Now()
	id,err := orm.NewOrm().Insert(user)
	return id,err
}

func TableName() (string){
	return "user"
}

func PageList(page,size int,filters map[string] interface{}) ([]*User,int64){
	users := make([]*User, 0)
	qt := orm.NewOrm().QueryTable(TableName())
	if len(filters) > 0 {
		for k,v := range filters{
			qt = qt.Filter(k,v)
		}
	}
	total,_:= qt.Count()
	qt.OrderBy("-id").Limit(size,(page-1)*size).All(&users)
	return users,total
}

func UpdateUser(user *User)  (int64,error) {

	o := orm.NewOrm()
	u := User{Id:user.Id}
	err :=  o.Read(&u)
	if err == nil{
		u.UserName = user.UserName
		u.Status = user.Status
		u.Remark = user.Remark
		u.Password = user.Password
		return o.Update(&u)
	}else {
		return 0,err
	}
}