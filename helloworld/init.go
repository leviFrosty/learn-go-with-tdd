package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
