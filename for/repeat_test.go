package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := strings.Repeat("a", 5)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("z", 10)
		expected := strings.Repeat("z", 10)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("r", 0)
		expected := strings.Repeat("r", 0)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
