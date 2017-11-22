package lib

import "time"

//日期格式化
func TimeFormat(tm time.Time)  (string){
	return tm.Format("2006-01-02 03:04:05")
}