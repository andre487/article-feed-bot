package main

import (
	ycSecretProvider "github.com/andre487/go-yc-secret-provider"
)

func GetBotToken(useProdToken bool) string {
	secretId := "e6qf87djrscl6bsvp102"
	if useProdToken {
		secretId = "e6qkeub22t7kqefji0sn"
	}
	val, err := ycSecretProvider.GetLockBoxTextValue(secretId, "token")
	PanicOnError(err)
	return val
}
