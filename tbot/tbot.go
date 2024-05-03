package tbot

import (
	"log"
	"os"
	"time"
	housemd "vzkguard/houseMD"

	tele "gopkg.in/telebot.v3"
)

var bot *tele.Bot
var userData *housemd.CacheUsers

func Start() {
	var err error
	log.Println("Starting VZKGuard...")
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
	log.Println("init user cache")
	userData = housemd.New()
	bot.Handle(tele.OnText, newMsg)
	log.Println("Bot Start")
	log.Println("----------------------")
	bot.Start()

}
