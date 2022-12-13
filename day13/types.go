package main

import (
	"aoc2022"
	"fmt"
	"strconv"
	"strings"
)

type ValueType int

const (
	IntValueType  ValueType = iota
	ListValueType ValueType = iota
)

type Value interface {
	AsList() List
	Compare(other Value) int
	String() string
	Type() ValueType
}

type List struct {
	data []Value
}

func (l List) AsList() List {
	return l
}

func (l List) Compare(other Value) int {
	return Pair{
		left:  l,
		right: other.AsList(),
	}.Compare()
}

func (l List) Iterate() *ListIterator {
	return &ListIterator{list: l}
}

func (l List) String() string {
	sb := strings.Builder{}

	sb.WriteString("[")

	for i, v := range l.data {
		if i > 0 {
			sb.WriteString(",")
		}

		sb.WriteString(" ")
		sb.WriteString(v.String())
	}

	sb.WriteString(" ]")

	return sb.String()
}

func (l List) Type() ValueType {
	return ListValueType
}

type IntValue struct {
	data int
}

func (i IntValue) AsList() List {
	return List{data: []Value{i}}
}

func (i IntValue) Compare(other Value) int {
	if other.Type() == IntValueType {
		return i.data - other.(IntValue).data
	}

	return Pair{
		left:  List{data: []Value{i}},
		right: other.(List),
	}.Compare()
}

func (i IntValue) String() string {
	return fmt.Sprintf("%d", i.data)
}

func (i IntValue) Type() ValueType {
	return IntValueType
}

func ParseList(s string) List {
	result, _ := parseList(s, 0)

	return result
}

func parseInt(s string, start int) (int, int) {
	for end := start + 1; end < len(s); end++ {
		if s[end] == ',' || s[end] == ']' {
			return aoc2022.Must(strconv.Atoi(s[start:end])), end - 1
		}
	}

	return -1, -1
}

func parseList(s string, start int) (List, int) {
	result := List{}
	i := start + 1

	for ; i < len(s); i++ {
		if s[i] == ',' {
			continue
		} else if s[i] == '[' {
			sl, j := parseList(s, i)

			result.data = append(result.data, sl)
			i = j
		} else if s[i] == ']' {
			break
		} else {
			value, j := parseInt(s, i)

			result.data = append(result.data, IntValue{data: value})
			i = j
		}
	}

	return result, i
}

type Pair struct {
	left  List
	right List
}

func (p Pair) Compare() int {
	li := p.left.Iterate()
	ri := p.right.Iterate()

	for li.HasNext() && ri.HasNext() {
		lv := li.Next()
		rv := ri.Next()
		result := lv.Compare(rv)

		if result != 0 {
			return result
		}
	}

	if li.HasNext() {
		return 1
	}

	if ri.HasNext() {
		return -1
	}

	return 0
}

func (p Pair) String() string {
	return fmt.Sprintf("Pair:\n  - left:  %v\n  - right: %v", p.left, p.right)
}

type ListIterator struct {
	list List
	pos  int
}

func (li *ListIterator) HasNext() bool {
	return li.pos < len(li.list.data)
}

func (li *ListIterator) Next() Value {
	result := li.list.data[li.pos]

	li.pos++

	return result
}

var DividerPackets = []List{
	ParseList("[[2]]"),
	ParseList("[[6]]"),
}
