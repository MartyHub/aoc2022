package main

import (
	"fmt"
	"log"
	"strings"
)

type Crate rune

type Stack struct {
	crates []Crate
}

func (s *Stack) String() string {
	return string(s.crates)
}

func (s *Stack) pop(n int) []Crate {
	l := len(s.crates)
	result := s.crates[l-n:]

	s.crates = s.crates[:l-n]

	return result
}

func (s *Stack) Push(crates []Crate) {
	for _, c := range crates {
		s.crates = append(s.crates, c)
	}
}

type Action struct {
	Amount int
	From   int
	To     int
}

func ParseAction(s string) Action {
	var a Action

	if _, err := fmt.Sscanf(s, "move %d from %d to %d", &a.Amount, &a.From, &a.To); err != nil {
		log.Fatal(err)
	}

	return a
}

func (a Action) String() string {
	return fmt.Sprintf("Move %d from %d to %d", a.Amount, a.From, a.To)
}

type Cargo struct {
	stacks []*Stack
}

func (c *Cargo) Add(s *Stack) {
	c.stacks = append(c.stacks, s)
}

func (c *Cargo) Move(a Action) {
	for i := 0; i < a.Amount; i++ {
		c.stacks[a.To-1].Push(c.stacks[a.From-1].pop(1))
	}
}

func (c *Cargo) Move2(a Action) {
	c.stacks[a.To-1].Push(c.stacks[a.From-1].pop(a.Amount))
}

func (c *Cargo) Top() string {
	sb := strings.Builder{}

	for _, s := range c.stacks {
		sb.WriteString(string(s.crates[len(s.crates)-1]))
	}

	return sb.String()
}
