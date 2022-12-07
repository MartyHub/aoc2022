package main

import (
	"aoc2022"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	lr := aoc2022.NewLineReader("input.txt")

	defer aoc2022.Close(lr)

	fs := &Dir{
		name:     "/",
		children: nil,
	}
	var pwd *Dir

	for lr.HasNext() {
		text := lr.Text()

		switch {
		case strings.HasPrefix(text, "$ cd"):
			dir := text[5:]

			switch dir {
			case "/":
				pwd = fs
			case "..":
				pwd = pwd.parent
			default:
				pwd = pwd.AddDir(dir)
			}
		case strings.HasPrefix(text, "$ ls"):
		case strings.HasPrefix(text, "dir "):
		default:
			name := ""
			size := 0

			if _, err := fmt.Sscanf(text, "%d %s", &size, &name); err != nil {
				log.Fatal(err)
			}

			pwd.AddFile(name, size)
		}
	}

	sum := 0

	fs.Do(func(d *Dir) {
		size := d.Size()

		if size <= 100000 {
			sum += size
		}
	})

	log.Printf("Sum: %d", sum) // 1583951

	total := 70000000
	unused := total - fs.Size()
	wanted := 30000000 - unused

	log.Printf("Wanted: %d", wanted)

	result := math.MaxInt

	fs.Do(func(d *Dir) {
		size := d.Size()

		if size >= wanted && size < result {
			result = size
		}
	})

	log.Printf("Result: %v", result)
}
