package main

func GCD(a, b int) int {
	for b != 0 {
		t := b

		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for _, v := range integers {
		result = LCM(result, v)
	}

	return result
}
