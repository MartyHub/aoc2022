package main

import "log"

func part2(monkeys *Monkeys) {
	log.Printf("Operation: %s", monkeys.byName["root"].Operation(monkeys))

	root := monkeys.byName["root"].(ComputingMonkey)
	expected := monkeys.byName[root.wait1].Yell(monkeys)
	guessMonkey := monkeys.byName[root.wait2]

	if !monkeys.byName[root.wait1].CanCompute(monkeys) {
		expected = monkeys.byName[root.wait2].Yell(monkeys)
		guessMonkey = monkeys.byName[root.wait1]
	}

	log.Printf("Expected: %d with %s", expected, guessMonkey.Operation(monkeys))
	log.Printf("Humn: %d", guessMonkey.(ComputingMonkey).ApplyOpposite(monkeys, expected)) // 3 352 886 133 831
}

func main() {
	testMonkeys := ParseMonkeys("test.txt")
	log.Printf("Test Root Monkey yells: %d", testMonkeys.Yell("root")) // 152

	monkeys := ParseMonkeys("input.txt")
	log.Printf("Root Monkey yells: %d", monkeys.Yell("root"))

	part2(monkeys)
}
