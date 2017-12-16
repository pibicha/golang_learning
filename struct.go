package main

import "fmt"

type Person struct {
	age  int32
	addr string
	name string
}

func main() {

	var P1, P2 Person
	P1.age = 10
	P1.addr = "北京"
	P1.name = "Monica"

	P2.age = 23
	P2.addr = "上海"
	P2.name = "Lulu"
	fmt.Println(P1, P2)

	var person_ptr *Person
	person_ptr = &P1
	fmt.Println(person_ptr)
}
