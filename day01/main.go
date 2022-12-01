package main

import (
	"log"
)

func main() {
	elves := MustRead("input.txt")

	log.Printf("Max Calories: %v", elves.MaxCalories()) // Elf # 189: 67 016 calories

	top3Calories := elves.TopCalories(3)

	log.Printf("Top 3 Calories: %v", top3Calories) // 200 116 calories

	top3Calories.Debug()
}
