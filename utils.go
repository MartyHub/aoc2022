package aoc2022

import (
	"io"
	"log"
	"strconv"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}

func PrettyFormat(i int) string {
	runes := []rune(strconv.Itoa(i))
	start := len(runes) - 1
	result := ""

	for i := start; i >= 0; i-- {
		if i != start && (start-i)%3 == 0 {
			result = " " + result
		}

		result = string(runes[i]) + result
	}

	return result
}
