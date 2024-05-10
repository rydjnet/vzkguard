package housemd

import (
	"log"
	"strings"
)

// HouseMD - Анализ на спам.  true - пользователю можно верить, fasle вероятно это спам
func HouseMD(tmsg *TMessage, cache *CacheUsers) bool {
	//Проверяем доверие к пользователю
	if cache.UserTrust(tmsg.User.Id) > 1 {
		log.Printf("User: %s (login: %s), have msgCount:%d\n", tmsg.User.UserFirstName, tmsg.User.UserLogin, cache.UserTrust(tmsg.User.Id))
		cache.NewMsg(tmsg.User)
		return true
	}

	log.Printf("User: %s (login: %s), have msgCount:%d\n", tmsg.User.UserFirstName, tmsg.User.UserLogin, cache.UserTrust(tmsg.User.Id))
	// Проверка на рекламу с верхним регистром
	if IsSpamToUpper(tmsg.Text) {
		return false
	}
	//Причесываем сообщение
	tmsg.Text = strings.ToLower(tmsg.Text)
	tmsg.Text = strings.ReplaceAll(tmsg.Text, ",", " ")
	tmsg.Text = strings.ReplaceAll(tmsg.Text, ".", " ")
	tmsg.Text = strings.ReplaceAll(tmsg.Text, "!", " ")
	tmsg.Text = strings.Join(strings.Fields(tmsg.Text), " ") //убираем лишние пробелы
	penaltyDicScore := 0
	//Проверяем через словарь
	penaltyDicScore += dicCheker(tmsg.Text)
	//Проверяем фразы
	penaltyPhasesScore := checkSpamPhrases(tmsg.Text)
	if penaltyPhasesScore > 0 && penaltyDicScore > 0 {
		return false
	}
	if penaltyDicScore > 2 {
		return false
	}
	// Если криминал не обнаружен, увеличиваем число сообщений.
	cache.NewMsg(tmsg.User)
	return true
}
