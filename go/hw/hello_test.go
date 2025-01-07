package main

import (
	"testing"
	"vexi.gg/utils"
)

func TestHello(t *testing.T) {
	t.Run("Say hello", func(t *testing.T) {
		got := Hello("vexi", "en")
		expected := "Hello vexi!"

		utils.AssertEqual(t, expected, got)
	})

	t.Run("Say hello with empty string", func(t *testing.T) {
		expected := "Hello world!"
		got := Hello("", "en")

		utils.AssertEqual(t, expected, got)
	})

	t.Run("Say hello in french", func(t *testing.T) {
		expected := "Bonjour vexi!"
		got := Hello("vexi", "fr")

		utils.AssertEqual(t, expected, got)
	})

	t.Run("Say hello in spanish", func(t *testing.T) {
		expected := "Hola vexi!"
		got := Hello("vexi", "sp")

		utils.AssertEqual(t, expected, got)
	})

	t.Run("Unknown language defaults to en", func(t *testing.T) {
		expected := "Hello vexi!"
		got := Hello("vexi", "unkown")

		utils.AssertEqual(t, expected, got)
	})
}
