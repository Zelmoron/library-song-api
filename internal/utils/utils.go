package utils

import "strings"

// splitIntoVerses разбивает текст песни на куплеты, каждый куплет — это строка.
func SplitIntoVerses(text string) []string {
	// Проверяем, не пустой ли текст
	if text == "" {
		return []string{}
	}

	// Разбиваем текст на строки
	lines := strings.Split(text, "\n")

	// Возвращаем результат — каждый элемент списка — это отдельная строка (куплет)
	return lines
}
