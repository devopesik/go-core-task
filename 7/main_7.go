package main

import (
	"fmt"
)

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, channel := range channels {
			for n := range channel {
				out <- n
			}
		}
	}()

	return out
}

func printChannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(c chan<- int) {
		defer close(c)
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}(ch1)

	go func(c chan<- int) {
		defer close(c)
		for i := -10; i <= 0; i++ {
			c <- i
		}
	}(ch2)

	ch := merge(ch1, ch2)
	printChannel(ch)

}
