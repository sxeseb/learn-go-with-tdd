package main

import "fmt"

const spanish = "SP"
const french = "FR"

const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "


func main() {
	fmt.Println(Hello("Chris", "SP"))
}

func Hello(n, l string) string {
	if n == "" {
		n = "world"
	}
	
	return resolvePrefix(l) + n
}

func resolvePrefix(l string) (p string) {
	switch l {
	case french:
		p = helloPrefixFrench
	case spanish:
		p = helloPrefixSpanish
	default:
		p = helloPrefixEnglish
	}

	return
}