package main

import "fmt"

func main() {

	//声明并初始化
	var s = "Let's go."
	fmt.Println(s)

	var b, c = 1, 100
	fmt.Println(b, c)

	//没有初始化就是默认值
	var d bool
	fmt.Println(d)

	f := "test"
	fmt.Println(f)

	var x, y = true, 2
	fmt.Println(x, y)

}
