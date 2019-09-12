package main

import "fmt"

func main() {
	fmt.Println("Simple string")
	fmt.Println(`
		This is a multi-line
		String, that can also contain "quotes"
`)
	fmt.Println("?")
	fmt.Println("\u23272")
	fmt.Println('L') // this is a rune
}
