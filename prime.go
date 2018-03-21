package main

import (
	"fmt"
)

func makeNatural() chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("makeNatural start.")
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func primeFilter(in <-chan int, prime int, count int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func doPrime() {

	ch := makeNatural()
	count := 1
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primeFilter(ch, prime, count)
		count++
	}
}
