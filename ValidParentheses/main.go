package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
    var ArrayParentheses = strings.Split(s, "")

		if len(ArrayParentheses) % 2 != 0 {
			return false
		}

		var stack []string
		for _, char := range ArrayParentheses {
			if char == "(" || char == "{" || char == "[" {
				stack = append(stack, char)
			} else {
				if len(stack) == 0 {
					return false
				}
				if char == ")" && stack[len(stack)-1] != "(" {
					return false
				}
				if char == "}" && stack[len(stack)-1] != "{" {
					return false
				}
				if char == "]" && stack[len(stack)-1] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			fmt.Println(stack)
		}

		return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[[])(]"))
}
