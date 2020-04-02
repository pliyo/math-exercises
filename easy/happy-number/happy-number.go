package main

func IsHappyNumber(number int) bool {
	uniqueSquares := make(map[int]int)
	uniqueSquares[number] = 1
	for {

		if number == 1 {
			return true
		}

		number = sumOfSquares(number)

		_, exists := uniqueSquares[number]
		if exists {
			return false
		}

		uniqueSquares[number] = 1
	}
	return false
}

func sumOfSquares(number int) int {
	var square = 0
	for number > 0 {
		var digit = number % 10
		square += digit * digit
		number = number / 10
	}
	return square
}

func main() {
}
