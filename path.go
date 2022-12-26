package aoc2022

import (
	"fmt"
	"strings"
)

type Step interface {
	comparable
}

type Path[S Step] struct {
	Steps []S
}

func NewPath[S Step](start S) Path[S] {
	return Path[S]{
		Steps: []S{start},
	}
}

func (p Path[S]) Clone() Path[S] {
	return Path[S]{
		Steps: append([]S{}, p.Steps...),
	}
}

func (p Path[S]) Extend(step S) Path[S] {
	l := p.Length()
	steps := make([]S, l+1)

	copy(steps, p.Steps)
	steps[l] = step

	return Path[S]{
		Steps: steps,
	}
}

func (p Path[S]) Contains(step S) bool {
	for _, s := range p.Steps {
		if s == step {
			return true
		}
	}

	return false
}

func (p Path[S]) First() S {
	return p.Steps[0]
}

func (p Path[S]) Last() S {
	return p.Steps[p.Length()-1]
}

func (p Path[S]) Length() int {
	return len(p.Steps)
}

func (p Path[S]) String() string {
	sb := strings.Builder{}

	sb.WriteString("[")

	for i, s := range p.Steps {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%v", s))
	}

	sb.WriteString("]")

	return sb.String()
}
