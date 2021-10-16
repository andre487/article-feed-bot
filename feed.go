package main

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func AddFeed(bot *tb.Bot, feedRequests chan *tb.Message) {
	for m := range feedRequests {
		fmt.Println(m.Payload)
		bot.Send(m.Sender, "Фид добавлен")
	}
}
