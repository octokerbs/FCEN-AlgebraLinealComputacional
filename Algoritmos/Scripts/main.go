package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	w := []int{4, 5, 6}

	fmt.Println(vecMult(v, w))
}

func vecMult(v, w []int) int {
	k := 0
	for i := range v {
		k += v[i] * w[i]
	}
	return k
}
