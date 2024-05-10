package tbot

import (
	"fmt"
	"log"
	"unicode/utf8"
	"vzkguard/antispam"
	"vzkguard/config"
	"vzkguard/perspective"

	tele "gopkg.in/telebot.v3"
)

// Создаем объект чата
var logChat = &tele.Chat{
	ID: int64(-1002036914981),
}

func msgHandler(m tele.Context) error {
	tmsg := m.Message()
	log.Printf("Group id: %d name: %s , UserName: %s, UserLogin: %s", tmsg.Chat.ID, tmsg.Chat.Title, tmsg.Sender.FirstName, tmsg.Sender.Username)
	_, ok := config.ChatsCfg[tmsg.Chat.ID]
	if !ok {
		return nil
	}
	if tmsg.Text == "" && tmsg.Caption == "" {
		return nil
	}
	if len(tmsg.Text) < 19000 || len(tmsg.Caption) < 19000 {
		Perspective(tmsg)
	}

	if utf8.RuneCountInString(tmsg.Text) < 500 || utf8.RuneCountInString(tmsg.Caption) < 500 {
		AntiSpam(tmsg)
	}
	return nil
}
func spam(m *tele.Message) {
	chatParams := config.ChatsCfg[m.Chat.ID]
	chatUser, _ := bot.ChatMemberOf(m.Chat, m.Sender)
	r := &tele.ReplyMarkup{}
	switch chatParams.BotMod {
	case config.ModeModer:
		alert := "Блокирую спам Пользователь: " + m.Sender.FirstName + " " + m.Sender.Username
		bot.Send(logChat, alert)
		bot.Forward(logChat, m)
		bot.Delete(m)
		bot.Ban(m.Chat, chatUser)

	case config.ModeWatcher:
		url := fmt.Sprintf("%s/%d", chatParams.Link, m.ID)
		btnReport := r.URL("Проверить", url)
		r.Inline(r.Row(btnReport))
		alert := "Вероятно обнаружен спамер: " + m.Sender.FirstName + " " + m.Sender.Username
		bot.Send(logChat, alert, r)
	}

}
func AntiSpam(m *tele.Message) {
	chatUser, _ := bot.ChatMemberOf(m.Chat, m.Sender)
	if chatUser.CanManageChat {
		return
	}
	if config.TUserCache.GetUser(m.Sender.ID) {
		return
	}
	var isSpam bool
	if m.Text != "" {
		isSpam = antispam.SpamDetecter(m.Text)
	} else {
		isSpam = antispam.SpamDetecter(m.Caption)
	}
	if isSpam {
		spam(m)
	}
	config.TUserCache.NewMsg(m.Sender.ID)
}

func Perspective(m *tele.Message) {
	toxicScore := getToxicScore(m)
	if toxicScore > float64(0.4) {
		chatParams := config.ChatsCfg[m.Chat.ID]
		r := &tele.ReplyMarkup{}
		url := fmt.Sprintf("%s/%d", chatParams.Link, m.ID)
		btnReport := r.URL("Проверить", url)
		r.Inline(r.Row(btnReport))
		alert := fmt.Sprintf("Сообщение пользователя %s оценивается в %.2f токсичности", m.Sender.FirstName, toxicScore)
		bot.Send(logChat, alert, r)
		bot.Forward(logChat, m)
	}
}

func getToxicScore(m *tele.Message) float64 {
	var toxicScore float64
	if m.Text != "" {
		toxicScore = perspective.ToxicCheker(m.Text)
	} else {
		toxicScore = perspective.ToxicCheker(m.Caption)
	}
	return toxicScore
}
