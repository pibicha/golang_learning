package main

import "fmt"

type Animal interface {
	utter()
}
type Dog struct {
}

/**
再定义一个方法，通过入参指定实现类，返回值指定实现方法；来完成继承
 */
func (hashiqi Dog) utter() {
	fmt.Println("fool people")
}

func main() {
	var animal Animal
	animal = new(Dog) //todo 多态需要使用new关键字？？？
	animal.utter()

	var hashiqi Dog // 不用new也可以创建，，
	hashiqi.utter()
}
