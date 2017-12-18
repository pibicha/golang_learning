package main

import (
	"fmt"
)

type People struct {
	age  int32
	addr string
	name string
}

func main() {

	// todo 这里导包有些问题，无法引用struct下的Person结构体
	name_person_map := make(map[string]People)

	var zhangsan, lisi, wangwu People
	zhangsan.age = 21
	zhangsan.name = "张三"
	zhangsan.addr = "北京"

	lisi.age = 25
	lisi.name = "李四"
	lisi.addr = "上海"

	wangwu.age = 31
	wangwu.name = "王五"
	wangwu.addr = "广州"

	name_person_map[zhangsan.name] = zhangsan
	name_person_map[lisi.name] = lisi
	name_person_map[wangwu.name] = wangwu

	delete(name_person_map, lisi.name)
	for name, person := range name_person_map {
		fmt.Println("性名：", name, "信息:", person)
	}

}
