package main

import (
	"aoc2022"
	"fmt"
	"log"
	"strconv"
)

const me = "humn"

type Monkey interface {
	ApplyOpposite(monkeys *Monkeys, n int) int
	CanCompute(monkeys *Monkeys) bool
	Name() string
	Operation(monkeys *Monkeys) string
	Yell(monkeys *Monkeys) int
}

type Monkeys struct {
	byName map[string]Monkey
	yells  map[string]int
}

func (m *Monkeys) Yell(name string) int {
	if n, found := m.yells[name]; found {
		return n
	}

	result := m.byName[name].Yell(m)

	m.yells[name] = result

	return result
}

func ParseMonkeys(fileName string) *Monkeys {
	lr := aoc2022.NewLineReader(fileName)

	defer aoc2022.Close(lr)

	result := &Monkeys{
		byName: map[string]Monkey{},
		yells:  map[string]int{},
	}

	for lr.HasNext() {
		m := parseMonkey(lr.Text())

		result.byName[m.Name()] = m
	}

	log.Printf("%d monkeys:", len(result.byName))

	for _, m := range result.byName {
		log.Printf(" - %v", m)
	}

	return result
}

func parseMonkey(s string) Monkey {
	snm := SpecificNumberMonkey{}

	if _, err := fmt.Sscanf(s, "%s %d", &snm.name, &snm.number); err == nil {
		snm.name = snm.name[:len(snm.name)-1]

		return snm
	}

	cm := ComputingMonkey{}

	if _, err := fmt.Sscanf(s, "%s %s %c %s", &cm.name, &cm.wait1, &cm.operation, &cm.wait2); err != nil {
		log.Fatalf("Failed to parse monkey '%s': %v", s, err)
	}

	cm.name = cm.name[:len(cm.name)-1]

	return cm
}

type SpecificNumberMonkey struct {
	name   string
	number int
}

func (m SpecificNumberMonkey) ApplyOpposite(_ *Monkeys, n int) int {
	if m.name == me {
		return n
	}

	panic(m.name)

	return 0
}

func (m SpecificNumberMonkey) CanCompute(_ *Monkeys) bool {
	return m.name != me
}

func (m SpecificNumberMonkey) Name() string {
	return m.name
}

func (m SpecificNumberMonkey) Operation(_ *Monkeys) string {
	if m.name == me {
		return "x"
	}

	return strconv.Itoa(m.number)
}

func (m SpecificNumberMonkey) Yell(_ *Monkeys) int {
	return m.number
}

func (m SpecificNumberMonkey) String() string {
	return fmt.Sprintf("%s: %d", m.name, m.number)
}

type ComputingMonkey struct {
	name         string
	wait1, wait2 string
	operation    rune
}

func (m ComputingMonkey) ApplyOpposite(monkeys *Monkeys, n int) int {
	monkey1 := monkeys.byName[m.wait1]
	monkey2 := monkeys.byName[m.wait2]

	switch m.operation {
	case '+':
		if monkey1.CanCompute(monkeys) {
			// n = monkey1.Yell() + XXX
			return monkey2.ApplyOpposite(monkeys, n-monkey1.Yell(monkeys))
		} else {
			// n = XXX + monkey2.Yell()
			return monkey1.ApplyOpposite(monkeys, n-monkey2.Yell(monkeys))
		}
	case '-':
		if monkey1.CanCompute(monkeys) {
			// n = monkey1.Yell() - XXX
			return monkey2.ApplyOpposite(monkeys, monkey1.Yell(monkeys)-n)
		} else {
			// n = XXX - monkey2.Yell()
			return monkey1.ApplyOpposite(monkeys, n+monkey2.Yell(monkeys))
		}
	case '*':
		if monkey1.CanCompute(monkeys) {
			// n = monkey1.Yell() * XXX
			return monkey2.ApplyOpposite(monkeys, n/monkey1.Yell(monkeys))
		} else {
			// n = XXX * monkey2.Yell()
			return monkey1.ApplyOpposite(monkeys, n/monkey2.Yell(monkeys))
		}
	case '/':
		if monkey1.CanCompute(monkeys) {
			// n = monkey1.Yell() / XXX
			return monkey2.ApplyOpposite(monkeys, monkey1.Yell(monkeys)/n)
		} else {
			// n = XXX / monkey2.Yell()
			return monkey1.ApplyOpposite(monkeys, n*monkey2.Yell(monkeys))
		}
	}

	log.Fatalf("Unknown operation '%c'", m.operation)

	return 0
}

func (m ComputingMonkey) CanCompute(monkeys *Monkeys) bool {
	return monkeys.byName[m.wait1].CanCompute(monkeys) && monkeys.byName[m.wait2].CanCompute(monkeys)
}

func (m ComputingMonkey) Name() string {
	return m.name
}

func (m ComputingMonkey) Operation(monkeys *Monkeys) string {
	operation1 := monkeys.byName[m.wait1].Operation(monkeys)
	operation2 := monkeys.byName[m.wait2].Operation(monkeys)

	if m.name == "root" {
		return fmt.Sprintf("%s == %s",
			operation1,
			operation2,
		)
	}

	switch m.operation {
	case '+':
		return fmt.Sprintf("%s + %s",
			operation1,
			operation2,
		)
	case '-':
		return fmt.Sprintf("%s - (%s)",
			operation1,
			operation2,
		)
	case '*':
		return fmt.Sprintf("(%s) * (%s)",
			operation1,
			operation2,
		)
	case '/':
		return fmt.Sprintf("(%s) / (%s)",
			operation1,
			operation2,
		)
	}

	log.Fatalf("Unknown operation '%c'", m.operation)

	return ""
}

func (m ComputingMonkey) Yell(monkeys *Monkeys) int {
	value1 := monkeys.Yell(m.wait1)
	value2 := monkeys.Yell(m.wait2)

	switch m.operation {
	case '+':
		return value1 + value2
	case '-':
		return value1 - value2
	case '*':
		return value1 * value2
	case '/':
		return value1 / value2
	}

	log.Fatalf("Unknown operation '%c'", m.operation)

	return 0
}

func (m ComputingMonkey) String() string {
	return fmt.Sprintf("%s: %s %c %s", m.name, m.wait1, m.operation, m.wait2)
}
