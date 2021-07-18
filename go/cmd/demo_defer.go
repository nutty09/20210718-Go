package main

import (
	"fmt"	"log"
)


func main() {
	fmt.Println("start")

	defer fmt.Println("TODO 0")
	// log.Fatal("Error na")
	for i := 0; i < 10; i++ {
		defer fmt.Println("TODO", i) // ย้อนกลับ
	}
	// defer fmt.Println("TODO")
	fmt.Println("finish")
}
