package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}

			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
