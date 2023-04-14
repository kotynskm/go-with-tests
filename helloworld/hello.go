package main

import (
	"fmt"
)
 const(
	helloPrefix = "Hello "
 	spanishPrefix = "Hola "
 	japanesePrefix = "Konnichiwa "
 	Spanish = "Spanish"
 	Japanese = "Japanese"
)

func sayHello(s, lang string) string {
	if s == "" {
		s = "World"
	}
	return greetingPrefix(lang) + s

}

func greetingPrefix(lang string) (prefix string){
	switch lang {
	case Spanish:
		prefix = spanishPrefix
	case Japanese:
		prefix = japanesePrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(sayHello("", ""))
	
}