package main

import "fmt"

const (
	englishPrefix = "Hello "
	frenchPrefix  = "Bonjour "
	spanishPrefix = "Hola "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
    prefix := englishPrefix
	switch lang {
	case "fr":
		prefix = frenchPrefix
	case "sp":
		prefix = spanishPrefix
	}
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", "en"))
}
