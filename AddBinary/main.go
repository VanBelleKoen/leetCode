package main

func addBinary(a string, b string) string {
	var result string
	var carry int
	for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		result = string(sum % 2 + '0') + result
		carry = sum / 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}

func main() {
	a := "11"
	b := "1"
	result := addBinary(a, b)
	println(result)
}