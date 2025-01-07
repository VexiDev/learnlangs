package utils

import "testing"

func AssertEqual[V comparable](t testing.TB, expected, actual V) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func AssertSlicesEqual[V comparable](t testing.TB, expected, actual []V) {
	t.Helper()
	len_a := len(actual)
	len_e := len(expected)
	if len_a != len_e {
		t.Errorf("Expected length %v, got %v", len_e, len_a)
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Value mismatch at index %d of %v: expected %v, found %v",
				i, actual, expected[i], actual[i])
		}
	}
}
