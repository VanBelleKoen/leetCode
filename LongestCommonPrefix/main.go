package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    var stringMatrix [][]string
		var commonPrefix string

		if len(strs) == 0 {
			return ""
		}
		
		for _, str := range strs {
			stringMatrix = append(stringMatrix, splitString(str))
		}

		for i := 0; i < len(stringMatrix[0]); i++ {
			for j := 1; j < len(stringMatrix); j++ {
				if i >= len(stringMatrix[j]) || stringMatrix[0][i] != stringMatrix[j][i] {
					return commonPrefix
				}
			}
			commonPrefix += stringMatrix[0][i]
		}
		return commonPrefix
}

func splitString(s string) []string {
	var result []string
	for _, char := range s {
		result = append(result, string(char))
	}
	return result
}

func main() {
		strs := []string{"dflower", "qflow", "cflight"}
		fmt.Println(longestCommonPrefix(strs))
}