package main

func main() {
}

//go:export add
func add(a, b int) int {
	return a + b
}

//go:export sub
func sub(a, b int) int {
	return a - b
}

//go:export mul
func mul(a, b int) int {
	return a * b
}

//go:export div
func div(a, b int) int {
	return a / b
}
