package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)
func main(){
	var (
		expr *cronexpr.Expression
		err error
		nowtime time.Time
		nexttime time.Time
	)
	//每分钟执行一次
	//分钟，小时，天，月，星期
	//支持秒级别 ，年
	//if expr,err := cronexpr.Parse("* * * * *"); err != nil{
	//	fmt.Println(expr,err)
	//	return
	//}

	//每隔5分钟执行1次
	if expr,err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(expr)

	}

	//当前时间
	nowtime = time.Now()
	//下次调度时间
	nexttime = expr.Next(nowtime)

	//等待这个定时器超时
	//time.NewTimer(nexttime.Sub(nowtime))

	time.AfterFunc(nexttime.Sub(nowtime), func() {
		fmt.Println("执行了",nexttime)
	})

	time.Sleep(5 * time.Second)
}
