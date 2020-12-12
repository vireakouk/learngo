package main

import (
	"fmt"
	"strings"
)

// shadow variable. n doesn't change
func shadowVariable() {
	n := 0
	if true {
		n := 1 // new variable is created
		n++
	}
	fmt.Println(n) // 0
}

func nonShadow() {
	n := 0
	if true {
		n = 1
		n++
	}
	fmt.Println(n) // 2
}

// In a multi-line slice, array or map literal, every line must end with a comma.
func expectingComma() {
	fruit := []string{"banana", "orange", "apple"} // no comma => ok
	vegetable := []string{
		"pepper",
		"cucumber",
		"mint", // no comma => error. Comma  required in multi-line slice/arragne/map
	}
	fmt.Println(fruit)
	fmt.Println(vegetable)
}

func immutableString() {
	s := "hello"
	// s[0] = 'H'
	fmt.Println(s) // cannot assign to s[0]. String is immutable
}

func mutableRune() {
	buf := []rune("hello")
	buf[0] = 'H' // single quote 'H' represents a single character (called a Rune)
	s := string(buf)
	fmt.Println(s)
}

func runeLiteral() {
	fmt.Println("H" + "i")                 // Hi
	fmt.Println('H' + 'i')                 // 177
	fmt.Println(string('H') + string('i')) // Hi

	s := fmt.Sprintf("%c%c, world!", 72, 'i')
	fmt.Println(s)
}

func abba() {
	fmt.Println(strings.TrimRight("Hello ABBA", "BA"))  // Trim the word
	fmt.Println(strings.TrimSuffix("Hello ABBA", "BA")) // Trim the matching string
}

func main() {
	abba()
}
