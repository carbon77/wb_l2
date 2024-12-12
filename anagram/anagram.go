package anagram

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки строки
func SortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Функция, проверяющая, являются ли две строки анаграммами
func CheckAnagrams(w1, w2 string) bool {
	return SortString(w1) == SortString(w2)
}

// Тип множества для хранения уникальных слов
type Set map[string]struct{}

// Функиця для поиска анаграмм
func FindAnagrams(words []string) map[string][]string {
	anagrams := make(map[string]Set)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		found := false
		for key := range anagrams {
			if CheckAnagrams(key, lowerWord) {
				anagrams[key][lowerWord] = struct{}{}
				found = true
				break
			}
		}

		// Если анаграмму не нашли, значит это первое вхождение, создаем новое множество
		if !found {
			anagrams[lowerWord] = Set{lowerWord: struct{}{}}
		}
	}

	result := map[string][]string{}
	for key, set := range anagrams {

		// Пропускаем множества с одним элементом
		if len(set) > 1 {

			// Конверитурем множество в слайс строк
			words = make([]string, 0, len(set))
			for word := range set {
				words = append(words, word)
			}

			// Сортируем слайс
			sort.Strings(words)
			result[key] = words
		}
	}

	return result
}

func Run() {
	words := []string{"листок", "пятак", "столик", "слиток", "пятка", "тяпка", "аткпя", "лук", "тяпка"}
	anagrams := FindAnagrams(words)

	for key, words := range anagrams {
		fmt.Printf("%s: %v\n", key, words)
	}
}
