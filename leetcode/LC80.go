package main

func main() {

}

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	p, q := 2, 2
	for ; q < length; {
		if nums[p-2] != nums[q] {
			nums[p] = nums[q]
			p++
		}
		q++
	}
	return p
}
