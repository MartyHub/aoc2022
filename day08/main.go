package main

import (
	"aoc2022"
	"log"
)

func read() [][]int {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	result := make([][]int, 0)

	for lr.HasNext() {
		text := lr.Text()
		row := make([]int, len(text))

		for i, c := range text {
			row[i] = int(c - '0')
		}

		result = append(result, row)
	}

	return result
}

func visibleColumn(x, y int, heights [][]int) bool {
	h := heights[y][x]
	invisible := 0

	for row := 0; row < y; row++ {
		if heights[row][x] >= h {
			invisible += 1

			break
		}
	}

	for row := y + 1; row < len(heights); row++ {
		if heights[row][x] >= h {
			invisible += 1

			break
		}
	}

	return invisible < 2
}

func visibleRow(x, y int, heights [][]int) bool {
	h := heights[y][x]
	invisible := 0

	for column := 0; column < x; column++ {
		if heights[y][column] >= h {
			invisible += 1

			break
		}
	}

	for column := x + 1; column < len(heights[0]); column++ {
		if heights[y][column] >= h {
			invisible += 1

			break
		}
	}

	return invisible < 2
}

func viewingDistanceTop(x, y int, heights [][]int) int {
	h := heights[y][x]
	result := 0

	for row := y - 1; row >= 0; row-- {
		result += 1

		if heights[row][x] >= h {
			break
		}
	}

	return result
}

func viewingDistanceBottom(x, y int, heights [][]int) int {
	h := heights[y][x]
	result := 0

	for row := y + 1; row < len(heights); row++ {
		result += 1

		if heights[row][x] >= h {
			break
		}
	}

	return result
}

func viewingDistanceLeft(x, y int, heights [][]int) int {
	h := heights[y][x]
	result := 0

	for column := x - 1; column >= 0; column-- {
		result += 1

		if heights[y][column] >= h {
			break
		}
	}

	return result
}

func viewingDistanceRight(x, y int, heights [][]int) int {
	h := heights[y][x]
	result := 0

	for column := x + 1; column < len(heights[0]); column++ {
		result += 1

		if heights[y][column] >= h {
			break
		}
	}

	return result
}

func viewingDistance(x, y int, heights [][]int) int {
	return viewingDistanceTop(x, y, heights) * viewingDistanceBottom(x, y, heights) * viewingDistanceLeft(x, y, heights) * viewingDistanceRight(x, y, heights)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	heights := read()
	height := len(heights)
	width := len(heights[0])
	visible := width*2 + (height-2)*2

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if visibleColumn(x, y, heights) || visibleRow(x, y, heights) {
				visible += 1
			}
		}
	}

	log.Printf("Visibles: %v", aoc2022.PrettyFormat(visible))

	scenicScore := 0

	for y, row := range heights {
		for x := range row {
			vd := viewingDistance(x, y, heights)

			scenicScore = max(scenicScore, vd)
		}
	}

	log.Printf("Scenic Score: %v", aoc2022.PrettyFormat(scenicScore))
}
