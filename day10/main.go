package main

import (
	"aoc2022"
	"log"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	cpu := Cpu()
	crt := Crt()
	cycle := 0
	sum := 0

	for lr.HasNext() {
		instruction := Parse(lr.Text())

		for !instruction.Complete() {
			cycle++

			if cycle == 20 || cycle > 20 && (cycle-20)%40 == 0 {
				signalStrength := cpu.X * cycle
				sum += signalStrength
			}

			crt.draw(cycle, cpu.X)
			instruction.Tick(cpu)
		}
	}

	log.Printf("Sum of Signal Strengths = %v", sum) // 10 760
	log.Printf("Crt:\n%v", crt.sb.String())         // F P G P H F G H
}
