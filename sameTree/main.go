package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			name:     "Both trees are nil",
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			name:     "One tree is nil",
			p:        &TreeNode{Val: 1},
			q:        nil,
			expected: false,
		},
		{
			name:     "Trees with different values",
			p:        &TreeNode{Val: 1},
			q:        &TreeNode{Val: 2},
			expected: false,
		},
		{
			name: "Identical trees",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "Different structures",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSameTree(tt.p, tt.q)
			if result != tt.expected {
				t.Errorf("isSameTree() = %v, expected %v", result, tt.expected)
			}
		})
	}
}