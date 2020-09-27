package main

import "fmt"

/**
 * To test, it's a good idea to separate complex constructs, into more atomic
 * ones. Hence, we separate fmt.Println and the String.
 */

const englishHelloPrefix = "Hello "

// Hello returns the String to be printed.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}