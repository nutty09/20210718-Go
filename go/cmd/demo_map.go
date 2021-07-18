package main

import "fmt"

func main() {

	n := 1
	switch n {
	case 1:
		fmt.Println("OK 1")
	case 2:
		fmt.Println("OK 2")
	case 3:
		fmt.Println("OK 3")
	default:
		fmt.Println("default")
	}

	switch {
	case n == 1:
		fmt.Println("OK 1")
	case n%2 == 0:
		fmt.Println("OK 2")
	case n > 10:
		fmt.Println("OK 3")
	default:
		fmt.Println("default")
	}

	// var grades map[int] string = {
	// 	1: "A"
	// }

	grades := make(map[int]string)
	grades[100] = "A"
	grades[101] = "B"
	grades[102] = "C+"

	for k, v := range grades {
		fmt.Println(k, v)
	}

	fmt.Println("Data of 100=%v", grades[100])
	v, found := grades[100]
	if found {
		fmt.Println("Found=", v)
	}

	if v, found := grades[100]; found {
		fmt.Println("Found=", v)
	}

	if v, found := grades[10011]; found {
		fmt.Println("Found=", v)
	} else {
		fmt.Println("Not found")
	}
}
