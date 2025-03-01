# Go スケジューラ (`robfig/cron` を使用)

このリポジトリでは、Go 言語のスケジューラライブラリ [`robfig/cron`](https://github.com/robfig/cron) を使用した **定期タスク実行プログラム** を提供します。

## **1. 概要**

このプロジェクトには 2 つのサンプルプログラムがあります。

### **1. `basic_cron.go`**

- **毎分** (`* * * * *`) 実行するジョブ
- **2 秒ごと (`@every 2s`)** に実行するジョブ
- **20 秒間だけ動作し、その後停止 (`c.Stop()`)**

### **2. `job_cron.go`**

- `cron.Job` インターフェースを実装した `MyJob` を使用
- **10 秒ごと (`*/10 * * * * *`)** に `MyJob.Run()` を実行
- プログラムを終了させずに **無限待機 (`select {}`)**

## **2. インストール**

### **Go のインストール**

このプログラムを動作させるには Go が必要です。  
[公式サイト](https://go.dev/doc/install) から Go をインストールしてください。

### **依存ライブラリのインストール**

このプロジェクトでは `robfig/cron` を使用します。以下のコマンドでインストールしてください。

```sh
go get github.com/robfig/cron/v3
```

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
