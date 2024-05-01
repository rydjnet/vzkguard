package housemd

import (
	"fmt"
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
		"$":              3,
	}
	msg = strings.ReplaceAll(msg, ",", "")
	arrWords := strings.Fields(msg)
	for _, val := range arrWords {
		penalty := pWords[val]
		penaltyScore += penalty
	}
	fmt.Println("Penalty: ", penaltyScore)
	return penaltyScore
}
