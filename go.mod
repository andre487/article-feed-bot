module github.com/andre487/article-feed-bot

go 1.17

replace github.com/andre487/go-yc-secret-provider => ./go-yc-secret-provider

require (
	github.com/andre487/go-yc-secret-provider v0.0.0-00010101000000-000000000000
	gopkg.in/tucnak/telebot.v2 v2.4.0
)

require github.com/pkg/errors v0.8.1 // indirect
