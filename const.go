package main

import (
	"unsafe"
	"fmt"
)

func main() {

	// 声明常量时，即使是作为局部变量也不会提示未被使用
	const b1 string = "permenent"
	const b2 string = "permenent"
	// const b2 := "permenet" 常量声明不能用:=
	const a1, a2, a3 = "1", "2", "3"

	// 枚举可以用常量实现~！ 在方法中就可以声明枚举~~
	const (
		c1 = "abc"
		c2 = len(c1) // len方法和python中一模一样。。。
		c3 = unsafe.Sizeof(c1)
	)
	fmt.Print(c1, c2, c3)

	// 使用枚举的时候，iota类似于枚举的元素下标，每次显示声明时即会赋予该元素所在位置的值(从0开始)
	const (
		a = iota // 如果第一个元素不使用iota 在往后的元素使用会报错
		b
		c
		d = "6"
		e
		f
		g = 100001
		h
		i = iota
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	// iota的值始终是对应元素下标的值，如果枚举使用了表达式(我把他理解为枚举的构造器)，那么所有的值必须是已赋值的；
	const (
		spring = 1<<iota + 10
		summer
		// summer = 128 << iota
		fall
		winter
	)
	fmt.Println(spring, summer, fall, winter)
}
