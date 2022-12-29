package aoc2022

func Copy[T any](s []T) []T {
	result := make([]T, len(s))

	copy(result, s)

	return result
}

func Contains[T comparable](s []T, value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func CopyAndAppend[T any](s []T, value T) []T {
	l := len(s)
	result := make([]T, l+1)

	copy(result, s)
	result[l] = value

	return result
}

func CopyAndAppends[T any](s []T, values []T) []T {
	l := len(s)
	result := make([]T, l+len(values))

	copy(result, s)

	return append(result, values...)
}

func Insert[T any](s []T, index int, value T) []T {
	if len(s) == index { // nil or empty slice or after last element
		return append(s, value)
	}

	s = append(s[:index+1], s[index:]...) // index < len(a)
	s[index] = value

	return s
}

func Prepend[T any](s []T, element T) []T {
	s = append(s, element)

	copy(s[1:], s)

	s[0] = element

	return s
}

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
