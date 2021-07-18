package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type processor interface {
	doSth()
	// doSth2()
}

type demo struct {
	id   int
	name string
}

func (d demo) doSth() {
	fmt.Println("call dosth")
}

// class demo implement Stringer {

// }

func (d demo) String() string {
	return fmt.Sprintf("id=%v, name=%v", d.id, d.name)
}

type A int

func (a A) doSth() {
	fmt.Println("call dosth from A")
}

func main() {
	// fmt.Printf("%+v", demo{})
	do(demo{})
	do(A(1))
}

// func do(d demo) {
// 	d.doSth()
// }

func do(p processor) {
	p.doSth()
}

// func do(p interface{}) {
// 	p.processor.doSth()
// }
