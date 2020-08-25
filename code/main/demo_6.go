package main

import "fmt"

func main() {

	var slice1 [] int = [] int{1, 2, 3}
	var slice2 = [] int{1, 2, 3}
	slice3 := [] string{"hi", "bro"}

	s := make([]int, 5, 10)
	printSlice(s)
	printSlice(slice1)
	printSlice(slice2)
	fmt.Println(len(slice3))
	s = append(s, 15)
	s = append(s, 25, 50)
	printSlice(s)

	s1 := make([]int, len(s), (cap(s))*2)
	copy(s1, s)
	printSlice(s1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
