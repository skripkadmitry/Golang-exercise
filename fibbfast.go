package main

import (
	"fmt"
	"math/rand"
)

func Fibbonachi(in <-chan int, out chan<- uint64) {
	found := make([]uint64, 2, 10)
	found[0], found[1] = 1, 1
	for {
		select {
		case x := <-in:
			if x < 0 {
				fmt.Println("Terminating...")
				return
			}
			for i := len(found); i <= x; i++ {
				found = append(found, found[len(found)-1]+found[len(found)-2])
			}
			out <- found[x]
		}
	}
}

func main() {
	inp := make(chan int)
	outp := make(chan uint64)
	defer close(inp)
	defer close(outp)
	go Fibbonachi(inp, outp)
	for j := 0; j < 1000; j++ {
		test := 500 + rand.Intn(j+100)
		inp <- test
		fmt.Printf("%dth fibbonachi number is %d\n", test, <-outp)
	}
	for j := 0; j < 100; j++ {
		inp <- j
		fmt.Printf("%dth fibbonachi number is %d\n", j, <-outp)
	}
	inp <- -1
}
