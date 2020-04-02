package main

func SingleNumber(nums []int) int {
	numbers := make(map[int]int)
	for _, number := range nums {
		numbers[number]++
	}

	for key, value := range numbers {
		if value == 1 {
			return key
		}
	}

	return -1
}

func main() {
}
