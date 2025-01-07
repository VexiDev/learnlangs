package main

import (
	"testing"
	"vexi.gg/utils"
)

func TestSum(t *testing.T) {
	t.Run("Sum of a slice numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		utils.AssertEqual(t, expected, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of 2 slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		got := SumAll(slice1, slice2)
		expected := []int{3, 9}

		utils.AssertSlicesEqual(t, expected, got)
	})

	t.Run("Sum of 1 slice", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		expected := []int{3}

		utils.AssertSlicesEqual(t, expected, got)
	})
}
