package config

import (
	"fmt"
	"log"
	"os"
	"vzkguard/cache"

	"gopkg.in/yaml.v2"
)

var ChatsCfg map[int64]ChatConfig
var PerspectiveToken string
var TUserCache cache.CacheUsers

func Init() {

	var chats []ChatConfig
	confFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	PerspectiveToken = os.Getenv("PERSPECT_TOKEN")
	if PerspectiveToken == "" {
		log.Fatalln("Perspective API Token(PERSPECT_TOKEN) indefined")
	}

	err = yaml.Unmarshal(confFile, &chats)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}
	fmt.Fprintln(os.Stdout, []any{chats}...)
	ChatsCfg = make(map[int64]ChatConfig)
	for _, chat := range chats {
		ChatsCfg[chat.ChatID] = chat
	}
	TUserCache = *cache.New()
}
