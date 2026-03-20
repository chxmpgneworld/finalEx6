package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var (
	ErrNoString = errors.New("have no string")
)

func Determinant(input string) (string, error) {
	if input == "" {
		return "", ErrNoString
	}
	var output string
	msymbols := ".- "
	if strings.Trim(input, msymbols) == "" {
		output = morse.ToText(input)
		return output, nil
	}
	output = morse.ToMorse(input)
	return output, nil
}
