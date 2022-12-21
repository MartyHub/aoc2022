package main

import (
	"aoc2022"
	"log"
	"reflect"
	"time"
)

func test() {
	e := ParseEncryptedFile("test.txt")
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, -3, 3, -2, 0, 4}) {
		log.Fatalf("Invalid parsing: %v", e)
	}

	e.Move(0)
	if !reflect.DeepEqual(e.Numbers(), []int{2, 1, -3, 3, -2, 0, 4}) {
		log.Fatalf("1 moves between 2 and -3: %v", e)
	}

	e.Move(0)
	if !reflect.DeepEqual(e.Numbers(), []int{1, -3, 2, 3, -2, 0, 4}) {
		log.Fatalf("2 moves between -3 and 3: %v", e)
	}

	e.Move(1)
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, 3, -2, -3, 0, 4}) {
		log.Fatalf("-3 moves between -2 and 0: %v", e)
	}

	e.Move(2)
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, -2, -3, 0, 3, 4}) {
		log.Fatalf("3 moves between 0 and 4: %v", e)
	}

	e.Move(2)
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, -3, 0, 3, 4, -2}) {
		log.Fatalf("-2 moves between 4 and 1: %v", e)
	}

	e.Move(3)
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, -3, 0, 3, 4, -2}) {
		log.Fatalf("0 does not move: %v", e)
	}

	e.Move(5)
	if !reflect.DeepEqual(e.Numbers(), []int{1, 2, -3, 4, 0, 3, -2}) {
		log.Fatalf("4 moves between -3 and 0: %v", e)
	}

	log.Printf("Decrypted: %v", e)
	log.Printf("Zero Index: %d", e.ZeroIndex())
	log.Printf("1000th: %d", e.Value((e.ZeroIndex()+1_000)%e.length).Number)
	log.Printf("2000th: %d", e.Value((e.ZeroIndex()+2_000)%e.length).Number)
	log.Printf("3000th: %d", e.Value((e.ZeroIndex()+3_000)%e.length).Number)
}

func main() {
	test()

	e1 := ParseEncryptedFile("input.txt")
	e1.Decrypt()
	log.Printf("Grove Coordinates Sum: %d", e1.GroveCoordinatesSum()) // 5 904

	e2 := ParseEncryptedFile("input.txt")
	e2.ApplyDecryptionKey()
	for i := 0; i < 10; i++ {
		start := time.Now()
		e2.Decrypt()
		log.Printf("Round # %d in %v", i+1, time.Since(start))
	}
	log.Printf("Grove Coordinates Sum: %d", e2.GroveCoordinatesSum())

	aoc2022.Print()
}
