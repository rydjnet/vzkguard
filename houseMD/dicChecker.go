package housemd

import (
	"fmt"
	"log"
	"strings"
)

var pWords = map[string]int{
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
	"гарантированно": 1,
	"сoтрудничествo": 5, // содржит латинские символы
	"дocтoйный":      5, // содржит латинские символы
	"сотрудничество": 1,
	"удалённое":      1,
	"удаленке":       1,
	"удаленном":      1,
	"пишите":         1,
	"лс":             1,
	"доход":          3,
	"дохода":         3,
	"доходом":        3,
	"заработка":      1,
	"oпытa":          5, // содржит латинские символы
	"обучeние":       5, // содржит латинские символы
	"поддeржка":      1, // содржит латинские символы
	"быстрoгo":       5, // содржит латинские символы
	"зapaбoтoк":      5, // содржит латинские символы
	"заработок":      3,
	"фaльшивые":      1,

	"$":   1,
	"18+": 1,
}
var whiteList = map[string]bool{
	"врача":        true,
	"няк":          true,
	"бк":           true,
	"врач":         true,
	"колит":        true,
	"пациенты":     true,
	"стул":         true,
	"туалет":       true,
	"исследования": true,
	"причину":      true,
	"болячка":      true,
	"болезнь":      true,
	"болячкой":     true,
	"боли":         true,
	"кишечник":     true,
	"заболеванием": true,
	"заболевание":  true,
	"лекарства":    true,
	"диарея":       true,
	"срк":          true,
	"диагноз":      true,
	"диагноза":     true,
}

func dicCheker(msg string) int {
	var penaltyScore int

	msg = strings.ReplaceAll(msg, ",", "")
	msg = strings.ToLower(msg)
	arrWords := strings.Fields(msg)
	for _, val := range arrWords {
		if whiteList[val] {
			return 0
		}
		penalty := pWords[val]
		penaltyScore += penalty
	}
	fmt.Println("Penalty: ", penaltyScore)
	return penaltyScore
}

func checkSpamPhrases(msg string) int {
	dicPhrases := []string{
		"нужны люди на удалённый заработок",
		"ищу партнеров в новый проект если интересно пишите в личные сообщения",
		"дocтoйный зapaбoтoк", //фраза с латинскими символами
		"бeз oпытa",           //фраза с латинскими символами
		"бeрём бeз oпытa",
		"предлагаю сотрудничество удалённо",
		"предлагаю сотрудничество",
		"сотрудничество удаленно",
		"сотрудничество удалённо",
		"всем совершеннолетним удалённое сотрудничество",
		"совершеннолетние люди на удалённое сотрудничество",
		"ответственных ребят с амбициями в новый проект",
		"удаленная деятельность",
		"сотрудничество на удаленке с хорошим доходом",
		"партнеров в команду для получения доп дохода",
		"250$ в день",
		"3000$ в месяц",
		"вариант заработка в удаленном формате",
		"проходят в любой бaнкомaт",
		"пoтенциaл дoхoдa",
		"людей на сотрудничество",
		"фaльшивые рубли",
		"поддержка для быстрoгo старта",
		"гарантированно заработаешь",
		"вариант заработка",
		"в новый проект",
	}
	msg = strings.ToLower(msg)
	msg = strings.ReplaceAll(msg, ",", "")
	msg = strings.ReplaceAll(msg, ".", "")
	msg = strings.ReplaceAll(msg, "!", "")
	cleanedText := strings.Join(strings.Fields(msg), " ") //убираем лишние пробелы
	penaltyScore := 0
	for _, phrase := range dicPhrases {
		if strings.Contains(cleanedText, phrase) {
			log.Printf("Found phrase: %s\n", phrase)
			penaltyScore++
		}
	}
	return penaltyScore
}
