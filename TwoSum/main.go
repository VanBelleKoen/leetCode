package main

func twoSum(nums []int, target int) []int {
    var i, j int
    found := false

    for i = 0; i < len(nums); i++ {
        for j = i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                found = true
                break
            }
        }
        if found {
            break
        }
    }

    return []int{i, j}
}