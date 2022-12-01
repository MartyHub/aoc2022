package aoc2022

import "testing"

func TestPrettyFormat(t *testing.T) {
	type args struct {
		i int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Simple", args: args{i: 1}, want: "1"},
		{name: "Thousand", args: args{i: 1234}, want: "1 234"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyFormat(tt.args.i); got != tt.want {
				t.Errorf("PrettyFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
