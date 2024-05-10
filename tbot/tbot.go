package tbot

import (
	"log"
	"os"
	"time"
	"vzkguard/config"

	tele "gopkg.in/telebot.v3"
)

var bot *tele.Bot

func Start() {
	var err error
	log.Println("Starting VZKGuard...")
	config.Init()
	pref := tele.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err = tele.NewBot(pref)
	if err != nil {
		log.Println("failed to create bot")
		log.Fatal(err)
		return
	}
	bot.Handle(tele.OnText, msgHandler)
	bot.Handle(tele.OnPhoto, msgHandler)
	log.Println("Bot started")
	log.Println("----------------------")
	bot.Start()

}
