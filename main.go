package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/andrii-mk/telegram-simple/clients/telegram"
	"github.com/andrii-mk/telegram-simple/consumer/event_consumer"
	"github.com/andrii-mk/telegram-simple/events/telegram"
	"github.com/andrii-mk/telegram-simple/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "./files_storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage", err)
	}

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), s)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

	// fetcher = fetcher.New()

	// rocessor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
