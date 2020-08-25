package main

import "fmt"

func main() {

	var a1 [5] int = [5]int{1, 2, 3, 4, 5}
	var a2 [3] int = [...]int{11, 22, 33}
	var a3 = [2]bool{true, false}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a3[1])

	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	for index, value := range a2 {
		fmt.Printf("Element[%d] = %d\n", index, value)
	}

	for _, value := range a2 {
		fmt.Println(value)
	}

}
