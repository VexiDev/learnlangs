package main

import (
	"testing"
	"vexi.gg/utils"
)

func TestHello(t *testing.T) {
	t.Run("Say hello",
		func(t *testing.T) {
			got := Hello("vexi")
			expected := "Hello vexi!"

			utils.AssertEquals(t, expected, got)
		})

	t.Run("Say hello with empty string",
		func(t *testing.T) {
			got := Hello("")
			expected := "Hello world!"

			utils.AssertEquals(t, expected, got)
		})
}
