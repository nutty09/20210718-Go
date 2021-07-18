package day01

import "fmt"

func doSth(input int) func(i int) int {
	// return func(x int) int { return input * 2 }
	// return do
	return func(x int) int { return input * x }
}

func do(x int) int {
	return x * 123
}

//private
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// //public
// func Hello() {

// }
