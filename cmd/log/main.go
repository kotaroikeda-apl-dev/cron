package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// ロガーを作成
	logger := log.New(os.Stdout, "cron: ", log.LstdFlags)

	// cron インスタンスを作成し、ロギングを有効化
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(logger)))

	// 5秒ごとに実行するジョブを追加
	c.AddFunc("@every 5s", func() {
		fmt.Println("5秒ごとに実行:", time.Now())
	})

	// cron スタート
	c.Start()

	// プログラムを終了させないように待機
	select {}
}
