package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"
const englishGreetingPrefix = "Hello, "
const frenchGreetingPrefix = "Bonjour, "
const japaneseGreetingPrefix = "Ohayou, "
const spanishGreetingPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){

	switch language {
	case french:
		prefix = frenchGreetingPrefix
	case japanese:
		prefix = japaneseGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}