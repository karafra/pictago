package utils

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("should be able to map int to int", func(t *testing.T) {
		input := []int{1, 2, 3, 4}

		result := Map(input, func(v int) int {
			return v * v
		})

		expected := []int{1, 4, 9, 16}

		if len(result) != len(expected) {
			t.Fatalf("expected length %d, got %d", len(expected), len(result))
		}

		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("index %d: expected %d, got %d", i, expected[i], result[i])
			}
		}
	})

	t.Run("should be able to map int to string", func(t *testing.T) {
		input := []int{1, 2, 3}

		result := Map(input, func(v int) string {
			return strconv.Itoa(v)
		})

		expected := []string{"1", "2", "3"}

		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("index %d: expected %s, got %s", i, expected[i], result[i])
			}
		}
	})

	t.Run("should be able to handle empty slice", func(t *testing.T) {
		input := []int{}

		result := Map(input, func(v int) int {
			return v * 2
		})

		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("should preserve order of elements", func(t *testing.T) {
		input := []int{3, 1, 2}

		result := Map(input, func(v int) int {
			return v
		})

		expected := []int{3, 1, 2}

		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("order not preserved at index %d", i)
			}
		}
	})
}
