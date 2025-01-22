package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    return dummy.Next
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " ")
        head = head.Next
    }
    fmt.Println()
}

func main() {
    list1 := &ListNode{Val: 1}
    list1.Next = &ListNode{Val: 2}
    list1.Next.Next = &ListNode{Val: 4}

    list2 := &ListNode{Val: 1}
    list2.Next = &ListNode{Val: 3}
    list2.Next.Next = &ListNode{Val: 4}

    mergedList := mergeTwoLists(list1, list2)

    printList(mergedList)
}