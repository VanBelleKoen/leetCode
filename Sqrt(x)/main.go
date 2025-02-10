package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left < right {
		mid := left + (right - left) / 2
		if mid <= x / mid && (mid + 1) > x / (mid + 1) {
			return mid
		} else if mid > x / mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	x := 4
	result := mySqrt(x)
	println(result)
}