package main

func main() {

}

func removeElement(nums []int, val int) int {
	p, q := 0, 0
	for ; q < len(nums); {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
		}
		q++
	}
	return len(nums) - p
}
