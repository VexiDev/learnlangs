package main

import (
	"fmt"
	"testing"

	"vexi.gg/utils"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		expected := "aaaaa"
		got := Repeat("a", 5)

		utils.AssertEquals(t, expected, got)
	})

	t.Run("repeat 'a' -5 times", func(t *testing.T) {
		expected := ""
		got := Repeat("a", -5)

		utils.AssertEquals(t, expected, got)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}
