package main

import "fmt"

func main()  {
	letters := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(search(letters, 5))
	
}

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		midpoint := (max + min)/2
		if (nums[midpoint] == target) {
			return midpoint
		} else if (nums[midpoint] > target) {
			max = midpoint - 1
		} else {
			min = midpoint + 1
		}
	}
	return -1
}