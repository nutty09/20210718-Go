package main

import (
	"day01"
	"fmt"
)

// type department struct {
// 	depName string
// }

// type person struct {
// 	id   int
// 	name string
// 	department
// }

// // func (p person) doSth() {
// func (p *person) doSth() { //receiver

// 	p.name = "Update Name"
// 	fmt.Println("Called doSth() with name=", p.name)

// }

type parent struct {
}

type child struct {
	parent
}

func (p parent) callParent() {
	fmt.Println("callParent")

}

//overline
func (c child) callParent() {
	c.parent.callParent() //call from parent
	fmt.Println("callParent from child")

}

func main() {
	c := child{}
	c.callParent()
	// d := day01.Department{"demo"}
	// p := day01.Person{Id: 1, Name: "somkiat", Department: d}
	p := day01.NewPerson(1, "somkiat", "demo")
	p.DoSth()
	fmt.Printf("%+v", p)
	fmt.Println(p.DepName)
	fmt.Println(p.Department.DepName)
}

// class person {
// 	//properties
// 	id, name

// 	//method
// 	doSth()
// }
