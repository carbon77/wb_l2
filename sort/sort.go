package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortOptions struct {
	numeric bool
	reverse bool
	unique  bool
	column  int
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func sortLines(lines []string, opts *SortOptions) []string {
	if opts.unique {
		lines = removeDuplicates(lines)
	}

	sort.SliceStable(lines, func(i, j int) bool {
		var a, b string

		if opts.column != 0 {
			a = getColumn(lines[i], opts.column)
			b = getColumn(lines[j], opts.column)
		} else {
			a = lines[i]
			b = lines[j]
		}

		if opts.numeric {
			aNum, aErr := strconv.ParseFloat(a, 64)
			bNum, bErr := strconv.ParseFloat(b, 64)
			if aErr == nil && bErr == nil {
				if opts.reverse {
					return aNum < bNum
				}
				return aNum > bNum
			}
		}

		if opts.reverse {
			return a > b
		}
		return a < b
	})

	return lines
}

func getColumn(line string, i int) string {
	fields := strings.Fields(line)
	if i <= 0 || i > len(fields) {
		return ""
	}
	return fields[i-1]
}

func removeDuplicates(lines []string) []string {
	set := make(map[string]bool)
	var result []string

	for _, line := range lines {
		if !set[line] {
			set[line] = true
			result = append(result, line)
		}
	}

	return result
}

func main() {
	reverse := flag.Bool("r", false, "Reverse sorting")
	unique := flag.Bool("u", false, "Unique values")
	numeric := flag.Bool("n", false, "Sort by numeric value")
	column := flag.Int("k", 0, "Column number to sort (1-based)")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("wrong args number. want=1, got=%d", len(args))
	}

	opts := &SortOptions{
		reverse: *reverse,
		unique:  *unique,
		numeric: *numeric,
		column:  *column,
	}
	filename := args[0]
	lines := readLines(filename)
	lines = sortLines(lines, opts)
	err := writeLines(filename, lines)
	if err != nil {
		panic(err)
	}
}
