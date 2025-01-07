package main

import (
	"testing"
	"vexi.gg/utils"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat 'a' 5 times",
		func(t *testing.T) {
			expected := "aaaaa"
			got := Repeat("a")

			utils.AssertEquals(t, expected, got)
		})

}
