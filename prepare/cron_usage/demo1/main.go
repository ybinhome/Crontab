package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// linux crontab: 那一分钟 (0-59), 那一小时 (0-23), 那一天 (1-31), 那一月 (0-12), 星期几 (0-6)
	// cronexpr：支持秒级调度，支持年级调度 (2018-2099)

	// 1. 创建任务表达式，每 5 分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 2. 当前时间
	now = time.Now()
	// 3. 下次调度时间
	nextTime = expr.Next(now)

	// 4. nextTime.Sub(now) 计算出下次任务执行时间和当前时间的时间差；使用 time.AfterFunc 方法实现下次时间点执行函数中的任务；
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了：", nextTime)
	})

	// 主协程休眠 5 秒，等待任务执行完毕
	time.Sleep(5 * time.Second)
}
