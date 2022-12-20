package main

import (
	"log"
	"reflect"
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

	e := ParseEncryptedFile("input.txt")
	e.Decrypt()
	log.Printf("Grove Coordinates Sum: %d", e.GroveCoordinatesSum()) // 5 904
}
