package config

type BotMode string

const (
	ModeModer   BotMode = "moderator"
	ModeWatcher BotMode = "watcher"
)

type ChatConfig struct {
	ChatName string  `yaml: "chat_name"`
	ChatID   int64   `yaml:"chat_id"`
	Link     string  `yaml:"chat_link"`
	BotMod   BotMode `yaml:"mode"`
}

type BotConfig struct {
	ChatsList []ChatConfig `yaml:"ChatsList"`
	Chats     map[int64]ChatConfig
}
