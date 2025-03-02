package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// シンプルなロガー
	logger := log.New(os.Stdout, "[CRON] ", log.LstdFlags)

	// cron のログをシンプルにする
	c := cron.New(cron.WithLogger(cron.PrintfLogger(logger)))

	// 5秒ごとに実行
	_, err := c.AddFunc("@every 5s", func() { fmt.Println("5秒ごとに実行:", time.Now()) })
	if err != nil {
		log.Fatalf("ジョブの登録に失敗: %v", err)
	} else {
		log.Println("[CRON] ジョブが登録されました: @every 5s")
	}

	// cron スタート
	log.Println("[CRON] スケジューラを開始します")
	c.Start()

	// プログラムを終了させないように待機
	select {}
}
