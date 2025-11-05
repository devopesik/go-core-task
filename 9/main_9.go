package main

import (
	"fmt"
	"math"
)

func CubeConverter(in <-chan uint8, out chan<- float64) {
	go func() {
		defer close(out)
		for val := range in {
			out <- math.Pow(float64(val), float64(3))
		}
	}()
}

func main() {

	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go CubeConverter(ch1, ch2)

	go func(c chan<- uint8) {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- uint8(i)
		}
	}(ch1)

	for val := range ch2 {
		fmt.Println(val)
	}

}
