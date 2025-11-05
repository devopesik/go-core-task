package main

import "fmt"

func difference(a, b []string) []string {
	out := make([]string, 0)

	seen := make(map[string]struct{})
	for _, v := range b {
		seen[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := seen[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	slice1 := []string{"apple", "banana", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	slice := difference(slice1, slice2)
	fmt.Println(slice)
}
