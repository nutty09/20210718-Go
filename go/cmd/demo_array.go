package main

import "fmt"

func main() {
	// n := [5]int{1, 2, 3, 4, 5} //array
	n := []int{1, 2, 3, 4, 5} //slice
	fmt.Println(n[2:3])
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	n = append(n, 6) //append ใช้กับ slice
	for i, v := range n {
		fmt.Println(i, v)
	}

	i := 0
	for i < len(n) {
		fmt.Println(n[i])
		i++
	}
}
