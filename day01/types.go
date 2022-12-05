package main

import (
	"aoc2022"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type Elf struct {
	Id       int
	Calories int
}

func (e Elf) String() string {
	return fmt.Sprintf("Elf # %v: %v calories", e.Id, aoc2022.PrettyFormat(e.Calories))
}

type Elves struct {
	data []Elf
}

func MustRead(fileName string) Elves {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := Elves{}
	calories := 0

	for lr.HasNext() {
		text := lr.Text()

		if text == "" {
			result.data = result.newElf(calories)

			calories = 0
		} else {
			calories += aoc2022.Must(strconv.Atoi(text))
		}
	}

	if calories != 0 {
		result.data = result.newElf(calories)
	}

	log.Printf("Read %v", result)

	return result
}

func (e Elves) Len() int {
	return len(e.data)
}

func (e Elves) Calories() int {
	result := 0

	for _, elf := range e.data {
		result += elf.Calories
	}

	return result
}

func (e Elves) Debug() {
	for _, elf := range e.data {
		log.Println(elf)
	}
}

func (e Elves) MaxCalories() Elf {
	if e.Len() == 0 {
		return Elf{}
	}

	return e.data[0]
}

func (e Elves) TopCalories(n int) Elves {
	end := n
	l := e.Len()

	if end <= 0 || l < end {
		end = l
	}

	return Elves{
		data: e.data[0:end],
	}
}

func (e Elves) String() string {
	return fmt.Sprintf("%v elves carrying %v calories", e.Len(), aoc2022.PrettyFormat(e.Calories()))
}

func (e Elves) newElf(calories int) []Elf {
	l := e.Len()
	i, _ := sort.Find(l, func(i int) int {
		return e.data[i].Calories - calories
	})

	result := append(e.data, Elf{})

	copy(result[i+1:], result[i:])

	result[i] = Elf{Id: l + 1, Calories: calories}

	return result
}
