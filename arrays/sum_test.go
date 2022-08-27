package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}

		got := Sum(numbers)
		want := 20

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of some slices", func(t *testing.T) {
		got := SumAll([]int{2, 4}, []int{6, 8})
		want := []int{6, 14}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 4, 6}, []int{8, 10, 12})
		want := []int{10, 22}

		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{2, 4, 6, 8})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{2, 4}, []int{6, 8})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{2, 4, 6}, []int{8, 10, 12})
	}
}

func ExampleSum() {
	repeated := Sum([]int{2, 4, 6, 8})
	fmt.Println(repeated)
	// Output: 20
}

func ExampleSumAll() {
	repeated := SumAll([]int{2, 4}, []int{6, 8})
	fmt.Println(repeated)
	// Output: [6 14]
}

func ExampleSumAllTails() {
	repeated := SumAllTails([]int{2, 4, 6}, []int{8, 10, 12})
	fmt.Println(repeated)
	// Output: [10 22]
}
