package main

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 == 0 {
			n /= 3
			continue
		}

		return false
	}
	return true
}

func main() {
	println(isPowerOfThree(-3))
}
