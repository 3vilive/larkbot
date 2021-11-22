# larkbot

## usage

```go
package main

import "github.com/3vilive/larkbot"

func main() {
	bot := larkbot.NewBotWithSecretToken(
		"https://open.feishu.cn/open-apis/bot/v2/hook/your-webhook",
		"your secret token",
	)
	if err := bot.SendTextMessage("测试消息"); err != nil {
		panic(err)
	}
}
```
