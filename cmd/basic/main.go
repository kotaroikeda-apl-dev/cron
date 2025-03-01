package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// 毎分実行されるジョブ
	c.AddFunc("* * * * *", func() { fmt.Println("毎分実行:", time.Now()) })

	// 5秒ごとに実行するジョブ (@every)
	c.AddFunc("@every 2s", func() { fmt.Println("5秒ごと:", time.Now()) })

	// cron スタート
	c.Start()

	// プログラムを終了させないために待機
	// select {}

	time.Sleep(20 * time.Second) // 20秒待機
	c.Stop()
}
