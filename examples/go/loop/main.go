package main

func main() {
}

//go:volatile
type vint int

var x vint

//go:export infinite_loop
func loop1() {
	x = 0
	for {
		x++
	}
}
