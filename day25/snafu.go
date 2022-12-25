package main

import (
	"math"
	"strings"
)

func toDecimal(s string) int {
	n := 0

	for _, c := range s {
		n = n * 5

		switch c {
		case '0':
		case '1':
			n++
		case '2':
			n += 2
		case '-':
			n--
		case '=':
			n -= 2
		}
	}

	return n
}

func toSnafu(n int) string {
	r := int(math.Remainder(float64(n), 5))
	i := n

	switch r {
	case 1:
		i--
	case 2:
		i -= 2
	case -1:
		i++
	case -2:
		i += 2
	}

	i = i / 5

	sb := strings.Builder{}

	switch i {
	case 0:
	case 1:
		sb.WriteRune('1')
	case 2:
		sb.WriteRune('2')
	default:
		sb.WriteString(toSnafu(i))
	}

	sb.WriteRune(formatRemainder(r))

	return sb.String()
}

func formatRemainder(r int) rune {
	switch r {
	case -1:
		return '-'
	case -2:
		return '='
	default:
		return rune(r + '0')
	}
}
