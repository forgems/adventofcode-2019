package main

import "fmt"

func isValid(n int) bool {
	digit := 10
	duplicate_len := 0
	has_valid_duplicate := false
	for ; n > 0; n /= 10 {
		new_digit := n % 10
		if new_digit > digit {
			return false
		}
		if new_digit == digit {
			duplicate_len += 1
		} else {
			if duplicate_len == 1 {
				has_valid_duplicate = true
			}
			duplicate_len = 0
		}
		digit = new_digit
	}
	return has_valid_duplicate || duplicate_len == 1
}

func main() {
	valid := 0
	fmt.Println(isValid(135669))
	fmt.Println(isValid(122345))
	fmt.Println(isValid(111122))
	fmt.Println(isValid(112233))
	fmt.Println(isValid(112345))

	fmt.Println(isValid(111111))
	fmt.Println(isValid(111123))
	fmt.Println(isValid(223450))
	fmt.Println(isValid(123789))
	fmt.Println(isValid(121111))
	fmt.Println(isValid(123444))

	for n := 235741; n <= 706948; n++ {
		if isValid(n) {
			// fmt.Println(n)
			valid += 1
		}
	}
	fmt.Println(valid)
}
