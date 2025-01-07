package utils

import "testing"

func AssertEquals[V comparable](t testing.TB, expected, actual V) {
    t.Helper()
    if actual != expected {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}
