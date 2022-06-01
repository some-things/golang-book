package main

import "fmt"

func main() {
	var x string = "Hello World"

	fmt.Println(x)

	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	x = "hello"
	var y string = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	const z string = "this is a const string"
	fmt.Println(z)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	// multiply by 2
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	// degrees F to C
	fmt.Print("Enter degrees in Fahrenheit: ")
	var degF float64
	fmt.Scanf("%f", &degF)

	degC := (degF - 32) * 5 / 9

	fmt.Println(degF, "degrees Fahrenheit is", degC, "degrees Celsius")

	// feet to meters
	fmt.Print("Enter feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048

	fmt.Println(feet, "feet is", meters, "meters")
}
