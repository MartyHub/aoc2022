package main

import "unicode"

type Supply rune

func (s Supply) Priority() int {
	if unicode.IsUpper(rune(s)) {
		return int(s) - int('A') + 27
	} else {
		return int(s) - int('a') + 1
	}
}

func (s Supply) String() string {
	return string(s)
}
