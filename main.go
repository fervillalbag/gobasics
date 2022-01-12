package main

import (
	"fmt"
)

// var x int
// var y string
// var z bool

type dinero int
var x dinero

func main() {
	// #############################################
	
	// x := 42
	// y := "James Bond"
	// z := true

	// fmt.Println(x, y, z)

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// #############################################

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// #############################################
	
	// Caracteres o símbolos de escape

	// \n   U+000A line feed or newline
	// \t   U+0009 horizontal tab
	// \\   U+005C backslash
	// \'   U+0027 single quote  (valid escape only within rune literals)
	// \"   U+0022 double quote  (valid escape only within string literals)

	// #############################################
	
	// x = 42
	// y = "James Bond"
	// z = true
	
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// s := fmt.Sprintf("%v %v %v", x, y, z)
	// fmt.Println(s)

	// #############################################

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	
	
	// https://go.dev/ref/spec
}