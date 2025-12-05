package main

import (
	"errors"
	"fmt"
)

var outSideMain string = "hello from out side main"

func greet() {
	fmt.Println("Welcome to Go")
}

func double(x int) int {
	return x * 2
}

func fullName(fname string, lname string) string {
	return fname + " " + lname
}

func data(name string, age int) (string, int) {
	return name, age
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		fmt.Println("Cannot divide by zero")
		return 0, errors.New("cant divide with zero")
	}

	return x / y, nil
}

func main() {
	// GO Packages
	// fmt
	// fmt.Println("Hello, World!")
	// math
	// fmt.Println("My fav num is:", rand.Intn(10))

	// Variables:
	//1. Declare an integer variable named score and print its zero value.

	// var score int
	// fmt.Println("1", score)

	//2. Create a variable city using short declaration and assign a city name.

	// city := "Germany"
	// fmt.Println("2", city)

	// 3. Declare two variables age (int) and name (string) in one line.

	// var age, name = 25, "Jaddu"
	// fmt.Println("3", age, name)

	// 4. Use type inference to create a variable isActive (bool).

	// isActive := true
	// var isActive = true
	// fmt.Println("4", isActive)

	// 5. What happens if you use := twice on the same variable? Try it.

	// isActive := false ---> Gives error
	// isActive = false //---> correct
	// fmt.Println("5", isActive)

	// 6. Create a constant Version with value "1.0.0".
	// const version = "1.0.0"
	// fmt.Println("6", version)

	// 7. Swap values of two variables a := 10 and b := 20.

	// a := 10
	// b := 20

	// a = b
	// b = a

	// a, b = b, a
	// fmt.Println("7. values after swapping:", a, b)
	// 8. Write a variable price as float and print it.

	// var proce float64 = 8.5
	// price := 8.5
	// fmt.Println("8", price)

	// 9. Create a variable without assigning a value, then print it (show zero value).

	// var message string
	// fmt.Println("9. Message:", message)

	// 10. Try creating package-level variable (outside main) and access it inside main().

	// fmt.Println("10", outSideMain)

	// FUNCTION
	// Write a function greet() that prints “Welcome to Go”.

	// greet()

	// Write a function double(x int) that returns x * 2.

	// double(2)
	// Write a function fullName(first, last string) that returns a single string "first last".

	// fullName("john", "doe")
	// Write a function that returns two values: name and age.

	// data("jaddu", 25)
	// divide(2, 0)
	// Write a function isPositive(n int) that returns true if number > 0.
	//
	// LOOPS
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(i)
	// }

	// for j := 0; j < 5; j ++ {
	// 	fmt.Println(j)
	// }
	// for j := 0; j < 5; j += 2 {
	// 	fmt.Println(j)
	// }
	// Print “GoLang” 5 times.
	// for i := 0; i <= 4; i++ {
	// 	fmt.Println("Golang")
	// }
	// Print even numbers from 2 to 20.
	// for i := 2; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	//
	// SWITCH
	// age := 17

	// switch {
	// case age < 18:
	// 	fmt.Println("You are a minor")
	// case age >= 18:
	// 	fmt.Println("You are an adult")
	// default:
	// 	fmt.Println("Invalid age")

	// }

}
