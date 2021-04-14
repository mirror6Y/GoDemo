package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
}

func moveZeroes(nums []int) {
	p, q := 0, 0
	for ; q < len(nums); {
		if nums[q] != 0 {
			nums[p] = nums[q]
			p++
		}
		q++
	}
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
	for _, a := range nums {
		fmt.Println(a)
	}
}
