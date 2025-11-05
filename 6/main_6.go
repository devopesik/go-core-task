package main

import (
	"fmt"
	"math/rand"
)

func generate(c chan<- int, size int, maxValue int) {
	defer close(c)
	for i := 0; i < size; i++ {
		generatedNumber := rand.Intn(maxValue)
		c <- generatedNumber
	}
}

func printChannel(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)

	go func(c chan<- int) {
		generate(c, 10, 100)
	}(ch)

	printChannel(ch)
}
