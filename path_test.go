package aoc2022

import (
	"fmt"
	"reflect"
	"testing"
)

type Cube struct {
	x, y, z int
}

func (c Cube) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.x, c.y, c.z)
}

func TestNewPath(t *testing.T) {
	type args[T Step] struct {
		start T
	}
	type testCase[T Step] struct {
		name string
		args args[T]
		want Path[T]
	}
	tests := []testCase[Cube]{
		{
			name: "NewPath",
			args: args[Cube]{start: Cube{x: 1, y: 2, z: 3}},
			want: Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPath(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Contains(t *testing.T) {
	type args[T Step] struct {
		step T
	}
	type testCase[T Step] struct {
		name string
		p    Path[T]
		args args[T]
		want bool
	}
	tests := []testCase[Cube]{
		{
			name: "Contains",
			p:    Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}}},
			args: args[Cube]{step: Cube{x: 1, y: 2, z: 3}},
			want: true,
		},
		{
			name: "DoesNotContains",
			p:    Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}}},
			args: args[Cube]{step: Cube{x: 1, y: 2, z: 4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Contains(tt.args.step); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Extend(t *testing.T) {
	type args[T Step] struct {
		step T
	}
	type testCase[T Step] struct {
		name string
		p    Path[T]
		args args[T]
		want Path[T]
	}
	tests := []testCase[Cube]{
		{
			name: "Extend",
			p:    Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}}},
			args: args[Cube]{step: Cube{x: 1, y: 2, z: 4}},
			want: Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}, {x: 1, y: 2, z: 4}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Extend(tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Last(t *testing.T) {
	type testCase[T Step] struct {
		name string
		p    Path[T]
		want T
	}
	tests := []testCase[Cube]{
		{
			name: "Last",
			p:    Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}, {x: 1, y: 2, z: 4}}},
			want: Cube{x: 1, y: 2, z: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Last(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_String(t *testing.T) {
	type testCase[T Step] struct {
		name string
		p    Path[T]
		want string
	}
	tests := []testCase[Cube]{
		{
			name: "String",
			p:    Path[Cube]{steps: []Cube{{x: 1, y: 2, z: 3}, {x: 1, y: 2, z: 4}}},
			want: "[(1, 2, 3), (1, 2, 4)]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
