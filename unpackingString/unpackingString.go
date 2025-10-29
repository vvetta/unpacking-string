package unpackingstring

import (
	"errors"
)

var (
	ErrTwoNumInRow = errors.New("Two num in row!")
)

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	sRunes := []rune(s)
	result := []rune{}

	lastEscCharIdx := -2
	lastChar := sRunes[0]
	for i, char := range sRunes {

		isCharInt := char >= 48 && char <= 57
		isCharEsc := char == 92
		isLastCharInt := lastChar >= 48 && lastChar <= 57
		isLastCharEsc := lastChar == 92
		isLastCharEscInt := lastEscCharIdx == i-1

		if isCharInt && !isLastCharEscInt && isLastCharInt {
			return "", ErrTwoNumInRow
		}

		if isCharEsc {
			lastChar = char
			continue
		}

		if isLastCharEsc && isCharInt {
			result = append(result, char)
			lastEscCharIdx = i
			lastChar = char
			continue
		}

		if !isCharInt && !isLastCharEsc && !isLastCharEscInt && !isCharEsc {
			// Добавляем простой символ, не число.
			result = append(result, char)
		}

		// Если char число и последний символ экранирован или он простой
		if isCharInt && (isLastCharEscInt || (!isLastCharInt && !isLastCharEsc)) {
			result = append(result, multChar(lastChar, char)...)
		}

		lastChar = char
	}

	if len(result) == 0 {
		return "", ErrTwoNumInRow
	}

	return string(result), nil
}

func multChar(char rune, number rune) []rune {
	var num int

	switch number {
	case 48:
		num = 0
	case 49:
		num = 1
	case 50:
		num = 2
	case 51:
		num = 3
	case 52:
		num = 4
	case 53:
		num = 5
	case 54:
		num = 6
	case 55:
		num = 7
	case 56:
		num = 8
	case 57:
		num = 9
	}

	if num == 0 {
		return []rune{}
	}

	result := []rune{}
	for range num - 1 {
		result = append(result, char)
	}

	return result
}
