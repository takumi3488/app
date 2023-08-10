# Notice Bot for gRPC

gRPCでの各種チャットアプリやSNSのBotへメッセージを送信するライブラリです。

## 対応

|アプリ|対応状況|
|---|---|
|LINE|✅|
|Discord|そのうち|
|Slack|気が向いたら|
|Mastodon|周りで流行ったら|
|Misskey|周りで流行ったら|

## 使い方

起動

```
docker run -d --restart unless-stopped -p 8080:8080 \
  -e LINE_TOKEN=<YOUR_LINE_TOKEN> \
  -e LINE_TO=<YOUR_LINE_ACCOUNT_ID> \
  takumi3488/notice_bot_for_grpc
```

メソッドは `pkg/grpc/notice.proto` を確認してください。
