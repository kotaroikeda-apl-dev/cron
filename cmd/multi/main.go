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
	c.AddFunc("@every 10s", func() { fmt.Println("ジョブ1: 10秒ごと", time.Now()) })

	// 15秒ごとに実行
	c.AddFunc("@every 15s", func() { fmt.Println("ジョブ2: 15秒ごと", time.Now()) })

	// 20秒ごとに実行
	c.AddFunc("@every 20s", func() { fmt.Println("ジョブ2: 20秒ごと", time.Now()) })

	// 30秒ごとに実行
	c.AddFunc("@every 30s", func() { fmt.Println("ジョブ3: 30秒ごと", time.Now()) })

	// 1分ごとに実行
	c.AddFunc("@every 1m", func() { fmt.Println("ジョブ4: 1分ごと", time.Now()) })

	// cron スタート
	c.Start()

	// プログラムを終了させないように待機
	select {}
}
