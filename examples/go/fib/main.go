package main

func main() {
}

//go:export fib
func fib(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

//go:export fib_loop
func fibLoop(n uint32) uint32 {
	var i, a, b, c uint32
	for b = 1; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
