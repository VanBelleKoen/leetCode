package main

import (
	"strings"
)

func romanToInt(s string) int {
	ListOfInts := translateString(s)
	var ToSum []int
	for i := 0; i < len(ListOfInts); i++ {
		if i+1 < len(ListOfInts) {
			if ListOfInts[i] < ListOfInts[i+1] {
				ToSum = append(ToSum, ListOfInts[i+1]-ListOfInts[i])
				i++
			} else {
				ToSum = append(ToSum, ListOfInts[i])
			}
		} else {
			ToSum = append(ToSum, ListOfInts[i])
		}
	}

	return sumWithForLoop(ToSum)
}

func sumWithForLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
			sum += num
	}
	return sum
}

func translateString(s string) []int {
	ListOfChars := strings.Split(s, "")
	var ListOfInts []int

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(ListOfChars); i++ {
		ListOfInts = append(ListOfInts, romanMap[ListOfChars[i]])
	}

	return ListOfInts
}

func main() {
	println(romanToInt("III"))     // 3
	println(romanToInt("IV"))      // 4
	println(romanToInt("IX"))      // 9
	println(romanToInt("LVIII"))   // 58
	println(romanToInt("MCMXCIV")) // 1994
}
