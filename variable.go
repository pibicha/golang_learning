package main

import "fmt"

/**
	问题1
	如果某个局部变量x没有被使用，会报 x declared and not used ；全局变量声明不使用也不会报错~~~（方法中声明常量时，也不会报错）
	问题2
	变量声明可以指定类型，也可以省略类型，或者直接使用:=的方式把var声明省略（这种方式和python类似，就用这种了~~ 哈哈哈）
	问题3
	todo 变量类型int32和int64存放的数据范围，会不会有溢出和强制转换的问题？
	问题4
    int32((1>>32)/2) 范围相当于java的int 、 int64((1>>64)/2)对应long， float32((1>>32)/2)对应float、float64((1>>64)/2)对应double、、自己算吧。。
	todo: 不过go的无符号uint(1>>32) ，在java中好像没有对应的吧？

 */

// 全局变量的声明方式 ；
var (
	global1 = true
	global2 = 1
)
var g3, g4 = 3, 4

// g5, g6 := 3, 1 这种python式的声明只能在作为局部变量时使用。。。

func main() {
	var a = 1.5
	var a1 float32 = 1.5
	var a2 float32 = 1.5
	fmt.Println(a, a1, a2)

	var b = 2
	var b1 int32 = 2
	var b2 int64 = 2
	fmt.Println(b, b1, b2)

	// 声明和python很像！ 类型自动推断，性能会比显示声明类型低么？
	c := 10
	c1 := 0.5
	fmt.Print(c, c1)

	// 多变量声明：(这里想写在另一个方法里面，不过还不会调用go的方法~~~)
	d1, d2, d3 := true, -1, "str"
	fmt.Println(d1, d2, d3)

	var e1, e2, e3 = 1, "I like go", 0.1
	fmt.Println(e1, e2, e3)

	// 同一类型的多个变量
	var f1, f2, f3 int32
	f1, f2, f3 = 1, 3, 5
	fmt.Println(f1, f2, f3)

	//fmt.Println(global1, global2)
}
