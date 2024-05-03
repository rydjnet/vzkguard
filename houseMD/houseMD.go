package housemd

import (
	"fmt"
	"log"
	"strings"
)

func Guard(msg string) int {
	var penaltyScore int
	pWords := map[string]int{
		"ищу":            1,
		"поиске":         1,
		"нужны":          1,
		"нyжен":          1,
		"предложение":    1,
		"партнера":       1,
		"партнеров":      1,
		"людей":          1,
		"люди":           1,
		"человек":        1,
		"ребят":          1,
		"сoтрудничествo": 1,
		"сотрудничество": 1,
		"удалённое":      1,
		"пишите":         1,
		"лс":             1,
		"доход":          3,
		"дохода":         3,
		"доходом":        3,
		"зapaбoтoк":      3,
		"заработок":      3,

		"$": 3,
	}
	msg = strings.ReplaceAll(msg, ",", "")
	msg = strings.ToLower(msg)
	arrWords := strings.Fields(msg)
	for _, val := range arrWords {
		penalty := pWords[val]
		penaltyScore += penalty
	}
	fmt.Println("Penalty: ", penaltyScore)
	return penaltyScore
}

// HouseMD - Анализ на спам.  true - пользователю можно верить, fasle вероятно это спам
func HouseMD(tmsg TMessage, cache *CacheUsers) bool {
	//Шаг первый проверяем доверие к пользователю
	if cache.UserTrust(tmsg.User.Id) > 4 {
		log.Printf("User: %s (login: %s), have msgCount:%d\n", tmsg.User.UserFirstName, tmsg.User.UserLogin, cache.UserTrust(tmsg.User.Id))
		cache.NewMsg(tmsg.User)
		return true
	}
	log.Printf("User: %s (login: %s), have msgCount:%d\n", tmsg.User.UserFirstName, tmsg.User.UserLogin, cache.UserTrust(tmsg.User.Id))
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
