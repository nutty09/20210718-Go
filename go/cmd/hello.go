package main

import (
	"day01"
	db "day01/db"
	"fmt"
)

type Day int //Alias type

func main() {

	var d Day
	d = 1

	x := 100 //assign ค่าครั้งแรก
	x = 123
	fmt.Println(x)
	fmt.Printf("Hello %d", x)

	result := day01.Hello("somkiat")
	fmt.Println(result)

	row, err := db.SaveData("error")
	// _, err := saveData("error")
	fmt.Println(row, err)
	if err != nil {
		fmt.Println(row, err)
	}
}
