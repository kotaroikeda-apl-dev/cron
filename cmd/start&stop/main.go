package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// 新しい cron インスタンスを作成
	c := cron.New()

	// 10秒ごとに実行
	c.AddFunc("@every 10s", func() { fmt.Println("10秒ごとに実行", time.Now()) })

	// cron スタート
	c.Start()

	// 30秒間だけ実行
	time.Sleep(30 * time.Second)

	// cron を停止
	fmt.Println("ジョブを停止")
	c.Stop()
}
