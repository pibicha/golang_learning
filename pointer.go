package main

import "fmt"

func main() {

	var ip *int

	intVal := 1

	ip = &intVal
	fmt.Println(ip)

}
