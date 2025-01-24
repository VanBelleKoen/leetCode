package main

import "fmt"

func removeElement(nums []int, val int) int {
    k := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}

func main() {
    // Example 1
    nums1 := []int{3, 2, 2, 3}
    val1 := 3
    k1 := removeElement(nums1, val1)
    fmt.Println("k:", k1)  // Expected: 2
    fmt.Println("nums:", nums1[:k1]) // Expected: [2, 2]

    // Example 2
    nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val2 := 2
    k2 := removeElement(nums2, val2)
    fmt.Println("k:", k2)  // Expected: 5
    fmt.Println("nums:", nums2[:k2]) // Expected: [0, 1, 4, 0, 3]
}
