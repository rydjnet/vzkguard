package antispam

import (
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

// checkLatinCharacters - Проверяет число латинских символов. Если больше одного возвращает true, если меньше то false.
func checkerLatinCharacters(word string) bool {
	count := 0
	for _, char := range word {
		if unicode.Is(unicode.Latin, char) {
			count++
		}
	}
	if len(word)-count < 1 {
		return false
	}
	return count > 1
}

// isSpamToUpper - ищет сообщения которые написаны с множественным переходом через строку и бОльшая часть символов написана в верхнем регистре.
func checkerToUpper(msg string) bool {
	count := strings.Count(msg, "\n\n")
	if count > 2 {
		msg = strings.ReplaceAll(msg, " ", "")
		msg = strings.ReplaceAll(msg, "\n\n", "")
		upperCounter := 0
		for _, s := range msg {
			if unicode.IsUpper(s) {
				upperCounter++
			}
		}
		percent := float64(upperCounter) / float64(utf8.RuneCountInString(msg)) * 100
		if int(percent) > 60 {
			return true
		}
	}
	return false
}

// checkerWords - Считает сколько слов триггеров найдено в сообщении
func checkerWords(arrWords []string) int {
	var penaltyScore int

	for _, val := range arrWords {
		penalty := triggerWords[val]
		penaltyScore += penalty
	}
	log.Println("Trigger words: ", penaltyScore)
	return penaltyScore
}

func checkErarn(msg *string) bool {
	if strings.Contains(*msg, "$") {
		for _, word := range triggerEarns {
			if strings.Contains(*msg, word) {
				return true
			}
		}
	}
	return false
}

// checkerWords - Функция считает сколько слов содержат латинские буквы, при условии что первый символ не относится к ASCII.
func checkerLatinWords(arrWords []string) int {
	var penaltyScore int
	// Если первая буква в слове не ASCII то вероятно это слово на русском или другом языке, начинаем искать в нем латиницу.
	for _, val := range arrWords {
		if val[0] > unicode.MaxASCII {
			if checkerLatinCharacters(val) {
				penaltyScore++
			}
		}
	}
	log.Println("latin words: ", penaltyScore)
	return penaltyScore
}
func checkerTelegramFolders(msg string) bool {
	count := strings.Count(msg, "https://t.me/addlist")
	return count > 3
}

func checkerException(arrWords []string) bool {
	for _, word := range arrWords {
		if whiteList[word] {
			return true
		}
	}
	return false
}
func checkerSpamPhrases(msg *string) int {
	penaltyScore := 0
	for _, phrase := range triggerPhrases {
		if strings.Contains(*msg, phrase) {
			penaltyScore++
		}
	}
	log.Println("spam Phrases: ", penaltyScore)
	return penaltyScore
}

func SpamDetecter(msg string) bool {
	if checkerToUpper(msg) {
		log.Println("Spam detected ToUpper>60%")
		return true
	}
	if checkErarn(&msg) {
		log.Println("Spam Detected $ plus earn word")
		return true
	}
	if checkerTelegramFolders(msg) {
		log.Println("Spam detected t folders")
		return true
	}
	msg = strings.ToLower(msg)
	msg = strings.ReplaceAll(msg, ",", " ")
	msg = strings.ReplaceAll(msg, ".", " ")
	msg = strings.ReplaceAll(msg, "!", " ")
	arrMsg := strings.Fields(msg)
	msg = strings.Join(arrMsg, " ")
	if checkerException(arrMsg) {
		return false
	}
	trigLatinWords := checkerLatinWords(arrMsg)
	if trigLatinWords > 3 {
		log.Println("Spam detected Words with latin symb>3")
		return true
	}
	trigWords := checkerWords(arrMsg)
	if trigWords > 2 {
		log.Println("Spam detected trig Words>2")
		return true
	}
	trigPhrases := checkerSpamPhrases(&msg)
	if trigPhrases > 0 && trigWords > 0 {
		log.Println("Spam detected trig phrases and trig words>0")
		return true
	}

	return false
}
