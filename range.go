package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	/**
	for x := range nums 相似于python的for x in range(nums)
	不同的是，每次遍历，x在go中对应nums元素的下标... 而python中是元素
	如果想在go中使用for ... range方式遍历切片，可以在解包时使用'_'变量，
	舍弃下标
	 */
	for _, num := range nums { // for ... range也适用数组
		// fmt.Println(num)
		sum += num
	}
	fmt.Println(sum)

	// 使用for range可以分解很多类型，字符串和map都可以分解；
	str := "abc123!@#"
	for index, ele := range (str) {
		fmt.Println(index, ele)
	}

}
