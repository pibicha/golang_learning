package main

import "fmt"

/**
	值和引用，与C++的类似;
	todo：使用go的数组实现杨辉三角
 */
func main() {
	a := 50
	b := "50"
	fmt.Println(&a, "\n", &b)

}
