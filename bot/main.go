package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	feedRequests := make(chan *tb.Message, 2048)

	bot, err := tb.NewBot(tb.Settings{
		Token:  GetBotToken(false),
		Poller: &tb.LongPoller{},
	})
	PanicOnError(err)

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Hello world!")
	})

	bot.Handle("/add", func(m *tb.Message) {
		feedRequests <- m
	})

	go AddFeed(bot, feedRequests)
	bot.Start()
}
