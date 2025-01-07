package main

import (
	"fmt"
	"testing"

	"vexi.gg/utils"
)

func TestAdd(t *testing.T) {
	t.Run("Add 1 and 2", func(t *testing.T) {
		expected := 3
		got := Add(1, 2)

		utils.AssertEquals(t, expected, got)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
