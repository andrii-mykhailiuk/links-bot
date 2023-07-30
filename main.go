package main

import (
	"flag"
	"log"

	"github.com/andrii-mk/telegram-simple/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// rocessor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
