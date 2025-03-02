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
	logger := log.New(os.Stdout, "[CRON] ", log.LstdFlags)

	// cron インスタンスを作成（シンプルなログを使用）
	c := cron.New(cron.WithLogger(cron.PrintfLogger(logger)))

	// 5秒ごとに実行するジョブを追加（エラーハンドリング付き）
	_, err := c.AddFunc("@every 5s", func() {
		fmt.Println("5秒ごとに実行:", time.Now())
	})
	if err != nil {
		log.Fatalf("[CRON] ジョブの登録に失敗: %v", err) // 致命的なエラーの場合、ログを出して終了
	} else {
		log.Println("[CRON] ジョブが登録されました: @every 5s")
	}

	// cron スタート（エラーハンドリング付き）
	log.Println("[CRON] スケジューラを開始します")
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("[CRON] スケジューラの実行中にエラーが発生: %v", r)
		}
	}()
	c.Start()

	// プログラムを終了させないように待機
	select {}
}
