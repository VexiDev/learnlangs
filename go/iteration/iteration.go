package main

func Repeat(text string, amount int) string {
	if amount <= 0 {
		return ""
	}
	var repeated string
	for i := 0; i < amount; i++ {
		repeated += text
	}
	return repeated
}
