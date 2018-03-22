package main

import (
	"context"
	"fmt"
)

func makeNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("makeNatural start.")
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("makeNatural stoped.")
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func primeFilter(ctx context.Context, in <-chan int, prime int, count int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("#%v Filter start.\n", count)
				select {
				case <-ctx.Done():
					fmt.Printf("#%v Filter stoped.\n", count)
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func doPrime() {

	ctx, cancel := context.WithCancel(context.Background())

	ch := makeNatural(ctx)
	count := 1
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primeFilter(ctx, ch, prime, count)
		count++
	}
	cancel()
}
