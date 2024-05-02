package tbot

import (
	"fmt"
	"log"
	housemd "vzkguard/houseMD"

	tele "gopkg.in/telebot.v3"
)

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

func newMsg(c tele.Context) error {
	log.Println("A new message received")
	message := c.Message()
	tuser := housemd.TUser{
		Id:            message.Sender.ID,
		UserLogin:     message.Sender.Username,
		UserFirstName: message.Sender.FirstName,
	}
	tmessage := housemd.TMessage{
		ID:   message.ID,
		User: tuser,
		Text: message.Text,
	}
	log.Printf("Group id: %d name: %s", message.Chat.ID, message.Chat.Title)
	log.Printf("Start HouseMD")
	if housemd.HouseMD(tmessage, userData) {
		return nil
	}
	log.Println("HouseMD found spam message")

	alert := "Вероятно обнаружен спаммер: " + message.Sender.FirstName + " " + message.Sender.Username

	r := &tele.ReplyMarkup{}
	if message.Chat.ID == -1001137424763 {
		url := fmt.Sprintf("https://t.me/nyak_bk_vzk/%d", message.ID)
		btnReport := r.URL("Проверить", url)
		r.Inline(r.Row(btnReport))
		admChatID := int64(-1001178620090)
		// Создаем объект чата
		chat := &tele.Chat{
			ID: admChatID,
		}
		bot.Send(chat, alert, r)

	} else {
		btnFalseReport := r.Data("Бан", "ban")
		r.Inline(r.Row(btnFalseReport))
		bot.Handle(&btnFalseReport, bReport)
	}
	chatID := int64(-4111795968) // Тестовый чат

	// Создаем объект чата
	chat := &tele.Chat{
		ID: chatID,
	}
	bot.Send(chat, alert, r)

	return nil
}
