package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Опции утилиты
type GrepOptions struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

// Функция для чтения файла
func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file.", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Фукнция для проверки строки на вхождение
func matchLine(pattern, line string, opt *GrepOptions) bool {
	if opt.IgnoreCase {
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}

	if opt.Fixed {
		return strings.Contains(line, pattern)
	}

	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		log.Fatal("Failed to match line.", err)
	}
	return matched
}

func main() {
	opt := &GrepOptions{}

	flag.IntVar(&opt.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&opt.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&opt.Context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&opt.Count, "c", false, "количество строк")
	flag.BoolVar(&opt.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&opt.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&opt.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&opt.LineNum, "n", false, "напечатать номер строки")

	flag.Parse()

	args := flag.Args()
	pattern, filename := args[0], args[1]

	lines := readLines(filename)
	// Слайс индексов подходящих строк
	var result []int
	var count int
	for i, line := range lines {
		match := matchLine(pattern, line, opt)

		if opt.Invert {
			match = !match
		}

		if match {
			count++
			result = append(result, i)
		}
	}

	if opt.Count {
		fmt.Println(count)
		return
	}

	if opt.Context > 0 {
		opt.After = opt.Context
		opt.Before = opt.Context
	}

	// Слайс индексов строк, которые необходимо напечатать
	var printResult []int
	for _, idx := range result {
		start := idx - opt.Before
		end := idx + opt.After + 1

		if start < 0 {
			start = 0
		} else if len(printResult) > 0 && printResult[len(printResult)-1] > start {
			start = printResult[len(printResult)-1] + 1
		}

		if end >= len(lines) {
			end = len(lines)
		}

		for i := start; i < end; i++ {
			printResult = append(printResult, i)
		}
	}

	for _, idx := range printResult {
		if opt.LineNum {
			fmt.Print(idx+1, ": ")
		}
		fmt.Println(lines[idx])
	}
}
