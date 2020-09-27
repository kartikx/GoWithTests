package main

import "fmt"

/**
 * To test, it's a good idea to separate complex constructs, into more atomic
 * ones. Hence, we separate fmt.Println and the String.
 */

const englishHelloPrefix = "Hello "

// Hello returns the String to be printed.
func Hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}