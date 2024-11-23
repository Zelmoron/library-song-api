package utils

import "strings"

// splitIntoVerses разбивает текст песни на куплеты
func SplitIntoVerses(text string) []string {
	// Проверяем, не пустой ли текст
	if text == "" {
		return []string{}
	}

	// Разбиваем текст на строки
	lines := strings.Split(text, "\n")

	return lines
}
