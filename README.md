# GoTGramBot

GoTGramBot is the Golang bindings for the Telegram Bot API

Most of the methods and types used in this library are auto-generated by scraping the official documentation of Telegram Bot Api

So there are high chances of getting bugs and other issues. So please let us know through the issue section about the bug you have encountered. If something isn't clear, open an issue or submit a pull request.

This lib is the successor of [TGramBot](https://github.com/KeralaBots/TGramBot) which is actually in python. But I ended up in much bigger problems in Tgrambot. Some of the functions are still not stable and are more complex to import it. So I thought of making a lib much similar to it in Go which is much better.

## Installing.

```
go get github.com/KeralaBots/GoTGramBot
```



_**This Library is still in its Alpha phase. More tests and analysis are still in progress to make it more user friendly and easy to use**_

## Example

```go
package main

import (
	"net/http"

	bot "github.com/KeralaBots/GoTGramBot"
	filters "github.com/KeralaBots/GoTGramBot/filters"
	types "github.com/KeralaBots/GoTGramBot/types"
)

func start(b *bot.Bot, m *types.Message) error {
	replyButton := [][]types.InlineKeyboardButton{
		{
			{
				Text:         "Hi",
				CallbackData: "test",
			},
		},
	}

	_, err := b.SendMessage(
		m.Chat.Id,
		"Hi",
		&bot.SendMessageOpts{ReplyMarkup: types.InlineKeyboardMarkup{InlineKeyboard: replyButton}},
	)

	return err
}


func main() {
	tbot, _ := bot.CreateBot("token", &bot.ClientOpts{
		Client: http.Client{},
	})
  
	d := tbot.NewDispatcher()
	d.AddMessageHandler(start, filters.All)
  
  
	d.Run()
}
```

More examples are in [samples](https://github.com/KeralaBots/GoTGramBot/tree/main/samples) directory

## Credits

Special thanks to [Paul Larsen
](https://github.com/PaulSonOfLars) for his libraries
