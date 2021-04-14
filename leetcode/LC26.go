package main

import "fmt"

func main() {
	array := []int{1, 1, 2}
	fmt.Println(removeDuplicates(array))
}

func removeDuplicates(nums []int) int {
	p, q := 0, 0
	for ; p+1 <= len(nums) && q < len(nums); {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}
	return p+1
}
