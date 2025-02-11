package main

import (
	"reflect"
	"testing"
)

func listToSlice(head *ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}

func sliceToList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    head := &ListNode{Val: nums[0]}
    current := head
    for _, num := range nums[1:] {
        current.Next = &ListNode{Val: num}
        current = current.Next
    }
    return head
}

func TestDeleteDuplicates(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {[]int{1, 1, 2}, []int{1, 2}},
        {[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
        {[]int{1, 1, 1}, []int{1}},
        {[]int{1, 2, 3}, []int{1, 2, 3}},
        {[]int{}, []int{}},
    }

    for _, test := range tests {
        inputList := sliceToList(test.input)
        expectedList := sliceToList(test.expected)
        resultList := deleteDuplicates(inputList)
        if !reflect.DeepEqual(listToSlice(resultList), listToSlice(expectedList)) {
            t.Errorf("deleteDuplicates(%v) = %v; want %v", test.input, listToSlice(resultList), test.expected)
        }
    }
}