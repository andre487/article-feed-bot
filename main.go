package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  GetBotToken(false),
		Poller: &tb.LongPoller{},
	})
	PanicOnError(err)

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Hello world!")
	})

	bot.Start()
}
