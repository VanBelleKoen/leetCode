package main

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   int
    }{
        {[]int{1, 3, 5, 6}, 5, 2},
        {[]int{1, 3, 5, 6}, 2, 1},
        {[]int{1, 3, 5, 6}, 7, 4},
        {[]int{1, 3, 5, 6}, 0, 0},
        {[]int{1}, 0, 0},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("nums=%v,target=%v", tt.nums, tt.target), func(t *testing.T) {
            got := searchInsert(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("searchInsert(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}