package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// 毎分42分に実行 (動作確認用)
	c.AddFunc("42 * * * *", func() { fmt.Println("毎分実行:", time.Now()) })

	//平日6時〜18時に1時間ごと42分に実行
	c.AddFunc("42 6-18 * * 1-5", func() { fmt.Println("平日6時〜18時に実行") })

	// cron スタート
	c.Start()

	// プログラムを終了させずに待機
	select {}
}
