# Golang SSE Example

Go で実装した Server Sent Events の mock です。 

## How to start

```
docker compose up

or

go run ./cmd/main.go
```

mock_client/index.html をブラウザで開く。

server は `localhost:3000` で立ち上がるように初期設定されており、mock_client も `:3000` を向いています。

## How to use

mock client の各種ボタンの挙動は以下です。

- start: [`new EventSource(source)`](https://developer.mozilla.org/ja/docs/Web/API/EventSource) で source を poll します。メッセージを受信次第、body に受信メッセージを書き込みます。
- stop: [`eventSource.close()`](https://developer.mozilla.org/ja/docs/Web/API/EventSource/close) を呼び出し、source を空にセットします。
- clear: body に書き込まれた受信メッセージを clear します。接続状態には手を加えません。
