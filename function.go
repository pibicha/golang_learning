package main

import "fmt"

/**
方法调用时，基本类型是复制值操作，引用类型是传递地址值；
go的闭包用起来声明太详细了，看起来很冗长。。
 */
func main() {

	n1, n2 := 6, 5
	max := max(n1, n2)
	fmt.Println(max)

	s1, s2 := "你好", "明天"
	// 这里传递是字符串，也是引用类型；可以用这个实验引用类型的方法调用
	swap(s1, s2)
	fmt.Println(s1, s2)

	//fmt.Println(&s1, &s2)
	//swapEnhence1(&s1, &s2)
	//fmt.Println(&s1, &s2)

	fmt.Println(&s1, &s2)
	swapEnhence2(&s1, &s2)
	fmt.Println(s1, s2)
}

func max(num1, num2 int) int {

	if (num1 > num2) {
		return num1
	}
	return num2
}

/*
和python一样，方法可以返回多个值，但是要在方法声明处写好要返回的值列表
 */
func swap(x, y string) (string, string) {
	return y, x
}

/*
这里我交换了x,y的地址，但是此时的x、y仍然是值复制的形式。。。 只不过复制的是地址而已。。。
 */
func swapEnhence1(x, y *string) {
	x, y = y, x
	fmt.Println(*x, *y)
}

/**
想要交换x和y，只在方法中替换复制过来的xy是不行的，因为不会对调用方法的xy有影响，
只能将x和y地址所指向的内容交换了，
 */
func swapEnhence2(x, y *string) {
	*x, *y = *y, *x
	//fmt.Println(*x, *y)
}
