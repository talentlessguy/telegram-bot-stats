# telegram-bot-stats

Go library for collecting usage stats of telegram bots.

It produces the following json file:

```jsonc
{
  "count": 2,
  "ids": ["123", "456"] // telegram user IDs
}
```

## Install

```sh
go get github.com/talentlessguy/telegram-bot-stats
```

## Examples

### [telebot](https://github.com/tucnak/telebot)

```go
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
	stats "github.com/talentlessguy/telegram-bot-stats"
)

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  "TOKEN",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		stats.AddUserToStat(m)
	})

	b.Handle("/getstat", func(m *tb.Message) {

		json := stats.ParseStatJSON()
		out := "User count: " + strconv.Itoa(json.Count)
		b.Send(m.Sender, out)
	})

	b.Start()
}
```
