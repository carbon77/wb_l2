package unpack

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	digits = []rune("0123456789")
)

func Unpack(input string) (string, error) {
	runes := []rune(input)
	var result string
	var escape bool

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\\' && !escape {
			if i == len(runes)-1 {
				return "", errors.New("invalid input: string not terminated")
			}
			escape = true
			continue
		}

		if slices.Contains(digits, r) && !escape {
			return "", errors.New("invalid input: two unescaped digits in a row")
		}

		count := 1
		if i < len(runes)-1 && slices.Contains(digits, runes[i+1]) {
			count, _ = strconv.Atoi(string(runes[i+1]))
			i++
		}

		result += strings.Repeat(string(r), count)
		escape = false
	}
	return result, nil
}

func main() {
	input := "abcd\\"
	str, err := Unpack(input)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", str)
}
