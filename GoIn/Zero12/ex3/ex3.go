package ex3

import (
	"errors"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	runes := []rune(str)
	if position < 0 || position >= len(runes) {
		return 0, errors.New("position out of range")
	}
	return runes[position], nil
}
