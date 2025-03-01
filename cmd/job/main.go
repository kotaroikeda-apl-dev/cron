package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// MyJob 構造体
type MyJob struct{}

func (j MyJob) Run() {
	fmt.Println("実行中:", time.Now())
}

func main() {
	// 秒単位のスケジュールを有効にする
	c := cron.New(cron.WithSeconds())

	// 10秒ごとに実行
	c.AddJob("*/10 * * * * *", MyJob{})

	// cron スタート
	c.Start()

	// プログラムを終了させないように待機
	select {}
}
