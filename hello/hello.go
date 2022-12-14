package main

import (
	"fmt"
)

const french = "French"
const spanish = "Spanish"
const portuguese = "Portuguese"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "
const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Johny", "Portuguese"))
}
