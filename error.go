package main

import "fmt"
import (
	"errors"
	"math"
)

/**
抛出异常的时候，返回值nil代表错误消息
 */
func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 500, errors.New("入参不能为负数")
	}
	return math.Sqrt(num), nil
}

func main() {
	f, e := sqrt(64)
	fmt.Println(f, e)
	sqrt, i := sqrt(-12)
	fmt.Println(sqrt, i)
}
