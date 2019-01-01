package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTIme time.Time
}

func main() {
	//需要有1个调度协程，定时检查所有的cron任务，谁过期了就调用谁

	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)

	scheduleTable = make(map[string]*CronJob)
	//当前时间
	now = time.Now()

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTIme: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/4 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTIme: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	go func() {

		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)

		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				//判断是否过期
				if cronJob.nextTIme.Before(now) || cronJob.nextTIme.Equal(now) {
					go func(jobName string) {
						fmt.Println("执行：", jobName)
					}(jobName)

					//计算下次调度时间
					cronJob.nextTIme = cronJob.expr.Next(now)
					fmt.Println(jobName, "，下次调度时间：", cronJob.nextTIme)
				}
			}
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
