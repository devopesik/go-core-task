package main

import (
	"fmt"
	"math/rand"
)

const (
	randEndNumber = 100
	sliceSize     = 10
)

func main() {
	originalSlice := generateRandomSlice(sliceSize, randEndNumber)
	fmt.Printf("originalSlice: %v\n", originalSlice)
	evenSlice := sliceExample(originalSlice)
	fmt.Printf("evenSlice: %v\n", evenSlice)
	addNumSlice := addElements(originalSlice, 1337)
	fmt.Printf("addNumSlice: %v\n", addNumSlice)
	copiedSlice := copySlice(originalSlice)
	fmt.Printf("copiedSlice: %v\n", copiedSlice)
	removeElementSlice := removeElement(originalSlice, 0)
	fmt.Printf("removeElementSlice: %v\n", removeElementSlice)
}

func generateRandomSlice(size, max int) []int {
	used := make(map[int]struct{})
	result := make([]int, 0, size)

	for len(result) < size {
		generatedNum := rand.Intn(max)
		if _, ok := used[generatedNum]; !ok {
			used[generatedNum] = struct{}{}
			result = append(result, generatedNum)
		}
	}
	return result
}

func sliceExample(slice []int) []int {
	out := make([]int, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			out = append(out, slice[i])
		}
	}

	return out
}

func addElements(slice []int, num int) []int {
	out := make([]int, 0, len(slice)+1)
	out = append(slice, num)
	return out
}

func copySlice(slice []int) []int {
	out := make([]int, len(slice))
	copy(out, slice)
	return out
}

func removeElement(slice []int, ind int) []int {
	out := make([]int, 0, len(slice)-1)
	out = append(slice[:ind], slice[ind+1:]...)
	return out
}
