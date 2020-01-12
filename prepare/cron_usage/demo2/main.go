package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

// 1. 定义任务的结构体
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time // expr.Next(now)
}

func main() {
	// 需要有一个调度协程，它定时检查所有的 Cron 任务，谁过期了就执行谁

	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob // KEY: 任务的名字
	)

	// 2. 创建一个任务表，存放创建好的任务
	scheduleTable = make(map[string]*CronJob)

	// 当前时间
	now = time.Now()

	// 3. 定义 2 个 cronjob，放入任务表中
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob

	// 4. 启动一个协程，定时扫描任务表中的任务，谁过期了就执行谁
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)

		// 定时检查下任务调度表
		for {
			now = time.Now()

			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动一个协程，执行这个任务
					go func(jobName string) {
						fmt.Println("执行：", jobName)
					}(jobName)

					// 计算下一次执行时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间：", cronJob.nextTime)
				}
			}

			// 睡眠 100 毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: // 在 100 毫秒可读，返回
			}
		}
	}()

	// 主协程休眠 100 秒，查看定时任务调度过程
	time.Sleep(100 * time.Second)
}
