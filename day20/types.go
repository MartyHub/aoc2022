package main

import (
	"aoc2022"
	"fmt"
	"strconv"
	"strings"
)

type Value struct {
	InitialOrder int
	Number       int
}

type EncryptedFile struct {
	length int
	values *[]Value
}

func (e EncryptedFile) Decrypt() {
	for i := 1; ; i++ {
		index := e.NextIndexToMove(i)

		if index == -1 {
			break
		}

		e.Move(index)
	}
}

func (e EncryptedFile) GroveCoordinatesSum() int {
	zeroIndex := e.ZeroIndex()

	return e.Value((zeroIndex+1_000)%e.length).Number +
		e.Value((zeroIndex+2_000)%e.length).Number +
		e.Value((zeroIndex+3_000)%e.length).Number
}

func (e EncryptedFile) Move(index int) {
	newIndex := index
	value := e.Value(index)

	if value.Number > 0 {
		newIndex = e.NewRightIndex(index)
	} else if value.Number < 0 {
		newIndex = e.NewLeftIndex(index)
	}

	if newIndex != index {
		*e.values = aoc2022.Remove(*e.values, index)

		*e.values = aoc2022.Insert(*e.values, newIndex, value)
	}

	//log.Printf("Iteration # %d: Moving %d from index %d to %d => %v", value.InitialOrder, value.Number, index, newIndex, e)
}

// -1 a b c
//    a b c -1
//    a b -1 c

// a -1 b c
// -1 a b c
// a b c -1

func (e EncryptedFile) NewLeftIndex(index int) int {
	newIndex := index
	value := e.Value(index)

	for i := 0; i > value.Number; i-- {
		newIndex--

		if newIndex == 0 {
			newIndex = e.length - 1
		} else if newIndex < 0 {
			newIndex = e.length - 2
		}
	}

	return newIndex
}

func (e EncryptedFile) NewRightIndex(index int) int {
	newIndex := index
	value := e.Value(index)

	for i := 0; i < value.Number; i++ {
		newIndex++

		if newIndex == e.length {
			newIndex = 1
		}
	}

	return newIndex
}

func (e EncryptedFile) NextIndexToMove(initialOrder int) int {
	for i, v := range *e.values {
		if v.InitialOrder == initialOrder {
			return i
		}
	}

	return -1
}

func (e EncryptedFile) Numbers() []int {
	result := make([]int, e.length)

	for i, v := range *e.values {
		result[i] = v.Number
	}

	return result
}

func (e EncryptedFile) String() string {
	sb := strings.Builder{}

	for i, v := range *e.values {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%d", v.Number))
	}

	return sb.String()
}

func (e EncryptedFile) Value(index int) Value {
	return (*e.values)[index]
}

func (e EncryptedFile) ZeroIndex() int {
	for i, v := range *e.values {
		if v.Number == 0 {
			return i
		}
	}

	return -1
}

func ParseEncryptedFile(fileName string) EncryptedFile {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	values := make([]Value, 0)

	for lr.HasNext() {
		value := Value{InitialOrder: lr.Count(), Number: aoc2022.Must(strconv.Atoi(lr.Text()))}

		values = append(values, value)
	}

	return EncryptedFile{
		length: len(values),
		values: &values,
	}
}
