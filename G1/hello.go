package main

import "fmt"

/**
 * To test, it's a good idea to separate complex constructs, into more atomic
 * ones. Hence, we separate fmt.Println and the String.
 */

const spanish = "Spanish"
const french  = "French"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix  = "Bonjour "

// Hello returns the Greeting.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name + "!"
}

// Starts with lowercase, it's a private function.
func greetingPrefix(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}