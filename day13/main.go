package main

import (
	"aoc2022"
	"log"
	"sort"
)

func ParsePairs() []Pair {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	result := make([]Pair, 0)
	index := 0

	for lr.HasNext() {
		text := lr.Text()

		if text != "" {
			l := ParseList(text)

			if index%2 == 0 {
				result = append(result, Pair{left: l})
			} else {
				result[len(result)-1].right = l
			}

			index++
		}
	}

	return result
}

func main() {
	sum := 0
	lists := append([]List(nil), DividerPackets...)

	for i, pair := range ParsePairs() {
		log.Printf("Pair # %v: %v", i+1, pair)

		c := pair.Compare()

		log.Printf("Pair # %v is in the right order: %v", i+1, c < 0)

		if c < 0 {
			sum += i + 1
		}

		lists = append(lists, pair.left, pair.right)
	}

	log.Printf("Sum = %v", sum) // 5 198

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Compare(lists[j]) < 0
	})

	key := 1

	for _, p := range DividerPackets {
		i, _ := sort.Find(len(lists), func(i int) int {
			return p.Compare(lists[i])
		})

		log.Printf("Divider packet %v is at position %v", p, i+1)

		key *= i + 1
	}

	log.Printf("Key = %v", key) // 22 344
}
