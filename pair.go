package aoc2022

type Pair[T any] struct {
	First  T
	Second T
}

func NewPair[T any](first T, second T) Pair[T] {
	return Pair[T]{
		First:  first,
		Second: second,
	}
}
