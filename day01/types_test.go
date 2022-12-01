package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElf_String(t *testing.T) {
	type fields struct {
		Id       int
		Calories int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "ToString", fields: fields{Id: 1, Calories: 1234}, want: "Elf # 1: 1 234 calories"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elf{
				Id:       tt.fields.Id,
				Calories: tt.fields.Calories,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElves(t *testing.T) {
	elves := MustRead("input.txt")

	assert.NotNil(t, elves)

	assert.Equal(t, 249, elves.Len())

	assert.Equal(t, 11782837, elves.Calories())

	assert.NotNil(t, elves.MaxCalories())
	assert.Equal(t, 189, elves.MaxCalories().Id)
	assert.Equal(t, 67016, elves.MaxCalories().Calories)

	assert.NotNil(t, elves.TopCalories(1))
	assert.NotNil(t, elves.TopCalories(1).MaxCalories())
	assert.Equal(t, 189, elves.TopCalories(1).MaxCalories().Id)
	assert.Equal(t, 67016, elves.TopCalories(1).MaxCalories().Calories)
}
