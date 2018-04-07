package main

import (
	"flag"
)

func main() {
	n := flag.Int("n", 1, "print message n times")
	flag.Parse()
	for i := 0; i < *n; i++ {
		println("Hello, world!")
	}
}
