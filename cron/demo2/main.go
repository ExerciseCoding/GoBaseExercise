package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	nexttTime time.Time
}
func main(){
	//需要有1个调度协程，它定时检查所有的Cron任务，谁过期了就执行谁


	var (
		cronJob *CronJob
		expr *cronexpr.Expression
		now time.Time
		scheduleTable map[string]*CronJob //key : 任务名字
	)

	scheduleTable = make(map[string]*CronJob)
	//当前时间
	now = time.Now()
	//1.定义两个cronjob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:      expr,
		nexttTime: expr.Next(now),
	}

	//任务注册到调度表
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:      expr,
		nexttTime: expr.Next(now),
	}

	//任务注册到调度表
	scheduleTable["job2"] = cronJob

	//启动调度协程
	go func(){
		var (
			jobName string
			cronJob *CronJob
			now time.Time
		)
		//定时检查任务调度表
		for{
			now = time.Now()
			for jobName,cronJob = range scheduleTable{
				//判断是否过期
				if cronJob.nexttTime.Before(now) || cronJob.nexttTime.Equal(now){
					//启动一个协程，执行该任务
					go func(jobName string){
						fmt.Println("执行:",jobName)
					}(jobName)

					//计算下次调度时间
					cronJob.nexttTime = cronJob.expr.Next(now)
					fmt.Println("下次执行时间:",cronJob.nexttTime)
				}

			}

			//睡眠100毫秒
			select {
			case <- time.NewTimer(100 * time.Millisecond).C: //100毫秒返回可读
				
			}

		}

	}()
	time.Sleep(30 * time.Second)
}