package main

import "fmt"

func existsInSlice[T comparable](val T, values []T) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}

	return false
}

func Print[T any](s []T) {
	for _, val := range s {
		fmt.Println(val)
	}
}


func main() {
	ls := []int{1, 2, 3}
	Print(ls)
}