package main

import (
	"aoc2022"
	"log"
	"strconv"
	"strings"
)

type Register struct {
	X int
}

func Cpu() *Register {
	return &Register{X: 1}
}

type Instruction interface {
	Complete() bool
	Tick(r *Register)
}

type AddXInstruction struct {
	amount int
	cycle  int
}

func (i *AddXInstruction) Complete() bool {
	return i.cycle == 2
}

func (i *AddXInstruction) Tick(r *Register) {
	i.cycle++

	if i.cycle == 2 {
		r.X += i.amount
	}
}

type NoopInstruction struct {
	cycle int
}

func (i *NoopInstruction) Complete() bool {
	return i.cycle == 1
}

func (i *NoopInstruction) Tick(_ *Register) {
	i.cycle++
}

func Parse(s string) Instruction {
	if s == "noop" {
		return &NoopInstruction{}
	} else if strings.HasPrefix(s, "addx ") {
		return &AddXInstruction{amount: aoc2022.Must(strconv.Atoi(s[5:]))}
	}

	log.Fatalf("Unknown instruction: %v", s)

	return nil
}

type Screen struct {
	sb strings.Builder
}

func Crt() *Screen {
	return &Screen{
		sb: strings.Builder{},
	}
}

func (s *Screen) draw(cycle, x int) {
	pixel := (cycle - 1) % 40

	if pixel >= x-1 && pixel <= x+1 {
		s.sb.WriteString("#")
	} else {
		s.sb.WriteString(".")
	}

	if pixel == 39 {
		s.sb.WriteString("\n")
	}
}
