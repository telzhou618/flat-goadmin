package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id int64	`pk:"auto"`
	CreateTime time.Time `orm:"type(datetime)"`
	LogTitle string `orm:"size(100)"`
	LogContent string `orm:"null,type(text)"`
	LogParams string `orm:"null,size(500)"`
	LogReqUrl string `orm:"size(255)"`
	LogReqMethod string `orm:"size(10)"`
	ClientIp string `orm:"size(50)"`
	LogType int
	User *User `orm:"rel(fk)"`
}

func LogTableName() (string){
	return "log"
}

func LogPageList(page,size int,filters map[string] interface{}) ([]*Log,int64){
	logs := make([]*Log, 0)
	qt := orm.NewOrm().QueryTable(LogTableName())
	if len(filters) > 0 {
		for k,v := range filters{
			qt = qt.Filter(k,v)
		}
	}
	total,_:= qt.Count()
	qt.OrderBy("-id").Limit(size,(page-1)*size).All(&logs)
	return logs,total
}