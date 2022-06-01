package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // char at index 1 (e) represented as byte (101)
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         //false

	// 1364067664
	fmt.Println(32132 * 42452)

	// true
	fmt.Println((true && false) || (false && true) || !(false && false))
}
