package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	cargoDefinition := make([]string, 0)

	for lr.HasNext() {
		text := lr.Text()

		if text == "" {
			break
		}

		cargoDefinition = append(cargoDefinition, text)
	}

	stackIndexes := make([]int, 0)

	for i, c := range cargoDefinition[len(cargoDefinition)-1] {
		if c != ' ' {
			stackIndexes = append(stackIndexes, i)
		}
	}

	cargoDefinition = cargoDefinition[:len(cargoDefinition)-1]
	stacks := make([]*Stack, len(stackIndexes))

	for i := len(cargoDefinition) - 1; i >= 0; i-- {
		stackDefinition := cargoDefinition[i]

		for i, index := range stackIndexes {
			if stacks[i] == nil {
				stacks[i] = &Stack{}
			}

			if stackDefinition[index] != ' ' {
				stacks[i].Push([]Crate{Crate(stackDefinition[index])})
			}
		}
	}

	for i, s := range stacks {
		log.Printf("Stack %d: %s", i, s)
	}

	cargo := &Cargo{stacks: stacks}

	for lr.HasNext() {
		text := lr.Text()
		action := ParseAction(text)

		//cargo.Move(action)
		cargo.Move2(action)
	}

	log.Printf("Top: %v", cargo.Top()) // ZWHVFWQWW, HZFZCCWWV
}
