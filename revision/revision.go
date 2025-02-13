package main

import (
	"fmt"
	"time"
)

// func add(x int, y int) int {
//     return x + y
// }

// func greet(name string) {
//     fmt.Println("Hello, " + name + "!")
// }


func main() {
	// fmt.Println("Hello go, I'm back!")

	//? 1.variables

	// Explicitly set variables type as String
	// var name string = "Golang"

	// Type Inferred variables
	// var name = "Golang"

	// var age = 24

	// var percentage = 95.6

	// var isAdult = true

	// fmt.Println(name, age, isAdult, percentage)

	//  const country = "India"

	//  capital := "New Delhi"

	//  fmt.Println(country, capital)

	// Print the type of the variable
	// fmt.Println("Type of name:", reflect.TypeOf(name))
	// fmt.Println("Type of age:", reflect.TypeOf(age))
	// fmt.Println("Type of isAdult:", reflect.TypeOf(isAdult))
	// fmt.Println("Type of percentage:", reflect.TypeOf(percentage))

	//? 2. Array

	// var marks[5] int

	// zeroed value INT : 0 ([0,0,0,0,0])
	// String : ""      (["","","","",""])
	// Boolean : false  ([false,false,false,false,false])

	// fmt.Println(len(marks))
	// marks[0] = 10
	// marks[1] = 20
	// marks[2] = 30
	// marks[3] = 40
	// marks[4] = 50
	// fmt.Println(marks)

	// countries := [3]string {"India", "Indonesia", "Australia"}

	// fmt.Println(countries)

	// fmt.Println(reflect.TypeOf(countries))

	// Looping over arrays

	// for i := 0 ; i<len(countries) ; i++ {
	// 	fmt.Println(countries[i])
	// }

	// for index, value := range countries {
	// 	fmt.Println(index, value)
	// }

	//? 3. Slices (Dynamic Array)

	// 3.1 Using Var keyword

	// var marks []int
	// var names [4]int

	// fmt.Println(marks) // []
	// fmt.Println(names) // [0, 0, 0, 0]

	// fmt.Println(marks==nil) // True

	// fmt.Println(reflect.TypeOf(marks)) // [] int
	// fmt.Println(reflect.TypeOf(names)) // [4] int

	// 3.2 Using make function

	// var marks = make([]int, 4)

	// fmt.Println(marks) //[0, 0, 0, 0]

	// var students = make([]int, 4, 10) //(min(length), max(capacity))
	// or
	// students := make([]int, 4, 10)  // Shorthand

	// fmt.Println(students)  //[0, 0, 0, 0]
	// fmt.Println(len(students)) // length = 4
	// fmt.Println(cap(students)) // capacity = 10

	// 3.3 You can create a slice from an existing array by specifying a range.

	// var arr = [2]int {1,2}
	// arr := [4]int {1,2,3,4}

	// slice :=arr[1:4]

	// fmt.Println(arr)
	// fmt.Println(slice)

	//? 4. Maps

	// 4.1 Using map literals

	// m := map[string]int{"age": 24, "marks": 900}

	// fmt.Println(m)

	// fmt.Println (m["marks"])

	// _, ok := m["age"]

	// if ok {
	// 	fmt.Println("all ok", ok)
	// } else{
	// 	fmt.Println("Not all ok")
	// }

	//? 5. Functions

  
	// result := add(2, 3)
	// fmt.Println(result) 


	// greet("John")

	//? 6. Pointers

   // name := "John"

  // fmt.Println(&name)
  
  //? 7. Struct
	//    type order struct {
	// 	id string
	// 	price float64
	// 	status string
	// 	creadtedAt time.Time
	//    }

	// Reciever type

	// func (o *order) statusUpdate(status string) {

	// 	o.status = status

	// }

  

  myOrder := order{
	id: "1",
	price: 69,
	status: "Dispached", 
	creadtedAt: time.Now(),
  }

  myOrder.statusUpdate("Delivered")

  fmt.Println(myOrder)
}

type order struct {
	id string
	price float64
	status string
	creadtedAt time.Time
   }

func (o *order) statusUpdate(status string) {

	o.status = status

   }