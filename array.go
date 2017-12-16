package main

import (
	"fmt"
)

func main() {

	// 数组声明时必须显示申明大小或者隐式在{}中赋值
	scores1 := [4]int{1, 2, 3, 4}
	scores2 := [...]int{1, 2, 3}

	fmt.Println(scores1, scores2)

	//=======先声明后赋值=======
	var n [10]int
	var i int
	for i = 0; i < 10; i++ {
		n[i] = i
		fmt.Println(n[i])
	}
	//========生成杨辉三角======
	var yh [10][10]int
	var row, col int
	yh[0][0] = 1
	for row = 1; row < 10; row++ {
		// 先把都为1的元素设置上去
		yh[row][0] = 1
		yh[row][col] = 1
		if row == 9 { // 最后一行不能为下一行设置值
			break
		}
		for col = 0; col < row; col++ {
			var rbelow = yh[row][col] + yh[row][col+1]
			yh[row+1][col+1] = rbelow
		}
	}

	for row = 0; row < 10; row++ {
		for col = 0; col < row; col++ {
			fmt.Print(yh[row][col])
			fmt.Print("\t")
		}
		fmt.Println()
	}

}

/**
1
1 1
1 2 1
1 3 3 1
....
 */
// todo：返回二维数组的时候，返回值不会定义o(╥﹏╥)o
//func yanghui(scale int) *int {
//	// 构建规模为scale的一个二维数组
//	var yh [scale][scale]int
//	var row, col int
//	yh[0][0] = 1
//	for row = 1; row < scale; row++ {
//		// 先把都为1的元素设置上去
//		yh[row][0] = 1
//		yh[row][col] = 1
//		if row == scale-1 {
//			break
//		}
//		for col = 0; col < row; col++ {
//			var rbelow = yh[row][col] + yh[row][col+1]
//			yh[row+1][col+1] = rbelow
//		}
//	}
//	return &yh
//}
