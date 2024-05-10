package tbot

import (
	"fmt"
	"log"
	"unicode/utf8"
	"vzkguard/antispam"
	"vzkguard/config"
	housemd "vzkguard/houseMD"
	"vzkguard/perspective"

	tele "gopkg.in/telebot.v3"
)

// Создаем объект чата
var logChat = &tele.Chat{
	ID: int64(-1002036914981),
}

func bReport(c tele.Context) error {
	callback := c.Callback()
	fmt.Printf("Callback received: Unique: %s, Data: %s\n", callback.Unique, callback.Data)

	if callback.Unique == "report_spam" {
		fmt.Println("Handling report spam")
		// Ваша логика удаления сообщения и т.д.

	}

	// Важно вызвать c.Respond() для отправки уведомления об обработке колбэка
	c.Respond()
	return nil
}
func preInit(c tele.Context) (*housemd.TUser, *housemd.TMessage) {
	log.Println("Starting preCheck")
	var tmsg housemd.TMessage
	if c.Message().Text == "" {
		tmsg.Text = c.Message().Caption
	} else {
		tmsg.Text = c.Message().Text
	}
	tuser := housemd.TUser{
		Id:            c.Message().Sender.ID,
		UserLogin:     c.Message().Sender.Username,
		UserFirstName: c.Message().Sender.FirstName,
	}
	tmsg.User = tuser
	tmsg.ID = c.Message().ID
	return &tuser, &tmsg
}
func newMsg(c tele.Context) error {
	log.Println("A new message received")
	tuser, tmessage := preInit(c)
	if tmessage.Text == "" {
		return nil
	}
	message := c.Message()
	log.Printf("Group id: %d name: %s , UserName: %s, UserLogin: %s", message.Chat.ID, message.Chat.Title, tuser.UserFirstName, tuser.UserLogin)
	chatUser, _ := bot.ChatMemberOf(message.Chat, message.Sender)
	if housemd.Gentleman(tmessage) > float64(0.4) {
		chatID := int64(-1002036914981) // Тестовый чат
		// Создаем объект чата
		chat := &tele.Chat{
			ID: chatID,
		}
		alert := "Сообщение пользователя может быть воспринято как оскорбление, пожалкйста проверьте " + message.Sender.FirstName + " " + message.Sender.Username
		bot.Send(chat, alert)
		bot.Forward(chat, message)
	}
	// сообщение больше 1000 симвлов не проверяем
	if utf8.RuneCountInString(message.Text) > 500 {
		return nil
	}
	log.Println("Member can Manage Chat: ", chatUser.CanManageChat)
	// Проверяем на админа

	if chatUser.CanManageChat {
		return nil
	}
	chatParams, ok := config.ChatsCfg[message.Chat.ID]
	if !ok {
		return nil
	}

	log.Printf("Start HouseMD")
	if housemd.HouseMD(tmessage, userData) {
		return nil
	}
	log.Println("HouseMD found spam message ", chatParams.BotMod == config.ModeWatcher)

	r := &tele.ReplyMarkup{}
	switch chatParams.BotMod {
	case config.ModeModer:
		chatID := int64(-1002036914981) // Тестовый чат
		// Создаем объект чата
		chat := &tele.Chat{
			ID: chatID,
		}
		alert := "Блокирую спам Пользователь: " + message.Sender.FirstName + " " + message.Sender.Username
		bot.Send(chat, alert)
		bot.Forward(chat, message)
		bot.Delete(message)
		bot.Ban(message.Chat, chatUser)

	case config.ModeWatcher:
		chatID := int64(-1002036914981) // Тестовый чат
		// Создаем объект чата
		chat := &tele.Chat{
			ID: chatID,
		}
		url := fmt.Sprintf("%s/%d", chatParams.Link, message.ID)
		btnReport := r.URL("Проверить", url)
		r.Inline(r.Row(btnReport))
		alert := "Вероятно обнаружен спамер: " + message.Sender.FirstName + " " + message.Sender.Username
		bot.Send(chat, alert, r)
	}

	return nil
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
