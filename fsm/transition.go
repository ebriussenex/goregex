package fsm

import (
	"errors"
	"strings"
)

type Predicate struct {
	AllowedChars    string
	DisallowedChars string
}

func (p Predicate) check(input rune) (bool, error) {
	if p.AllowedChars != "" && p.DisallowedChars != "" {
		return false, errors.New("must be mutually exclusive")
	}

	if len(p.AllowedChars) > 0 {
		return strings.ContainsRune(p.AllowedChars, input), nil
	}

	if len(p.DisallowedChars) > 0 {
		return !strings.ContainsRune(p.DisallowedChars, input), nil
	}

	return false, nil
}

func (p Predicate) mustCheck(input rune) bool {
	res, err := p.check(input)
	if err != nil {
		panic(err.Error())
	}
	return res
}

type Transition struct {
	debugSymbol string
	to          *State
	from        *State
	predicate   Predicate
}
