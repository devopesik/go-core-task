package main

import "fmt"

func intercept(a, b []int) (bool, []int) {
	outMap := make(map[int]struct{})
	out := make([]int, 0)
	seen := make(map[int]struct{})

	for _, v := range a {
		seen[v] = struct{}{}
	}

	for _, v := range b {
		if _, exists := seen[v]; exists {
			if _, added := outMap[v]; !added {
				out = append(out, v)
				outMap[v] = struct{}{}
			}
		}
	}

	if len(out) == 0 {
		return false, out
	}

	return true, out

}

func main() {
	a := []int{65, 3, 58, 678, 64, 64}
	b := []int{64, 2, 3, 43, 64}
	ok, res := intercept(a, b)
	fmt.Println(ok, res)
}
