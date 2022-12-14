package main

import "fmt"

// Sum returns the sum of all numbers in a slice.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// SumAll returns the sum of all numbers in each slice.
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// SumAllTails returns the sum of all numbers in each slice except the first.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(SumAll([]int{1, 2}, []int{0, 9}))
	fmt.Println(SumAllTails([]int{1, 2}, []int{0, 9}))
}
