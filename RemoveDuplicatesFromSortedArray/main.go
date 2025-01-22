package main

import "fmt"

func removeDuplicates(nums []int) int {
	lastUnique := 0

	for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] { 
					nums[lastUnique + 1] = nums[i]
					lastUnique++ 
			}
	}

	return lastUnique + 1
}

func main() {
		nums := []int{1, 1, 2,}
		fmt.Println(removeDuplicates(nums))
}