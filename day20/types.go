package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const decryptionKey = 811589153

type CacheKey struct {
	Index  int
	Number int
}

type Value struct {
	InitialOrder int
	Number       int
}

type EncryptedFile struct {
	length int
	cache  map[CacheKey]int
	values []Value
}

func (e *EncryptedFile) ApplyDecryptionKey() {
	for i, v := range e.values {

		e.values[i].Number = v.Number * decryptionKey
	}
}

func (e *EncryptedFile) Decrypt() {
	for i := 1; ; i++ {
		index := e.NextIndexToMove(i)

		if index == -1 {
			break
		}

		e.Move(index)
	}
}

func (e *EncryptedFile) GroveCoordinatesSum() int {
	zeroIndex := e.ZeroIndex()

	return e.Value((zeroIndex+1_000)%e.length).Number +
		e.Value((zeroIndex+2_000)%e.length).Number +
		e.Value((zeroIndex+3_000)%e.length).Number
}

func (e *EncryptedFile) Move(index int) {
	value := e.Value(index)
	newIndex := e.newIndex(index, value)

	if newIndex != index {
		e.doMove(value, index, newIndex)
	}
}

func (e *EncryptedFile) doMove(value Value, fromIndex, toIndex int) {
	defer aoc2022.Timer("doMove")()

	e.values = aoc2022.Remove(e.values, fromIndex)
	e.values = aoc2022.Insert(e.values, toIndex, value)
}

func (e *EncryptedFile) newIndex(index int, value Value) int {
	key := CacheKey{Index: index, Number: value.Number}

	if newIndex, found := e.cache[key]; found {
		return newIndex
	}

	newIndex := index

	if value.Number > 0 {
		newIndex = e.newRightIndex(index, value)
	} else if value.Number < 0 {
		newIndex = e.newLeftIndex(index, value)
	}

	e.cache[key] = newIndex

	return newIndex
}

func (e *EncryptedFile) newLeftIndex(index int, value Value) int {
	defer aoc2022.Timer("newLeftIndex")()

	newIndex := (index + value.Number) % (e.length - 1)

	if newIndex <= 0 {
		newIndex += e.length - 1
	}

	return newIndex
}

func (e *EncryptedFile) newRightIndex(index int, value Value) int {
	defer aoc2022.Timer("newRightIndex")()

	return (index + value.Number) % (e.length - 1)
}

func (e *EncryptedFile) NextIndexToMove(initialOrder int) int {
	for i, v := range e.values {
		if v.InitialOrder == initialOrder {
			return i
		}
	}

	return -1
}

func (e *EncryptedFile) Numbers() []int {
	result := make([]int, e.length)

	for i, v := range e.values {
		result[i] = v.Number
	}

	return result
}

func (e *EncryptedFile) String() string {
	sb := strings.Builder{}

	for i, v := range e.values {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%d", v.Number))
	}

	return sb.String()
}

func (e *EncryptedFile) Value(index int) Value {
	return e.values[index]
}

func (e *EncryptedFile) ZeroIndex() int {
	for i, v := range e.values {
		if v.Number == 0 {
			return i
		}
	}

	return -1
}

func ParseEncryptedFile(fileName string) *EncryptedFile {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := &EncryptedFile{
		cache: map[CacheKey]int{},
	}

	for lr.HasNext() {
		value := Value{InitialOrder: lr.Count(), Number: aoc2022.Must(strconv.Atoi(lr.Text()))}

		result.values = append(result.values, value)
	}

	result.length = len(result.values)

	log.Printf("Length: %d", result.length)

	return result
}
