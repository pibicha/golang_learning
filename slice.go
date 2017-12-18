package main

import "fmt"

func main() {
	// 使用make创建切片 cap可选
	var slice1 []byte = make([]byte, 3) // 切片声明需要指定长度，容量可选；
	var slice2 []byte = make([]byte, 3, 5)
	fmt.Println(slice1, slice2)

	s1 := []int{1, 2, 3}  // 声明切片不需要指定长度,且这里cap=len
	s2 := [4]int{1, 2, 3} // 声明数组需要指定长度

	fmt.Println(s1, s2)

	// 切片的操作，与python类似

	// 拷贝切片
	s1_copy1 := s1[:]
	// s1_copy1 := s1[0:-1] 没有0到-1这种操作。。。
	fmt.Println(s1_copy1)

	// 不指定start则默认从0开始 ；不指定end会一直到最后一个元素
	s1_copy2 := s1[:3]
	fmt.Println(s1_copy2)

	// append方法，类似python的append
	s1 = append(s1, 4) // s1 := append(s1, 4) 这种用法是错误的额，:=只能用于创建新变量
	fmt.Println(s1)

	// 添加的字符时不会报错，是因为转成ascii了....
	s1 = append(s1, 4, 'a', 'b')
	// append(s1, "a","b")
	fmt.Println(s1)

	// todo：go语言的切片没有pop之类的方法么？

}
