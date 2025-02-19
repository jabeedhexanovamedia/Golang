# GoLang [![Go](https://img.shields.io/badge/Golang-v1.22.3-bred.svg)](https://golang.org/)

### Why Learn Go?

#### 1. Simplicity and Ease of Learning

- Go has a clean and concise syntax that is easy to understand, even for beginners.
- The language design emphasizes simplicity and readability, reducing the cognitive load for developers.

#### 2. Performance

- Go is a statically typed, compiled language that offers performance close to that of lower-level languages like C and C++.
- It produces fast and efficient executables, which makes it suitable for high-performance applications.

#### 3. Concurrency

- Go's concurrency model is one of its standout features. Goroutines and channels provide an efficient way to handle concurrent tasks.
- This makes Go particularly well-suited for building scalable and high-concurrency applications, such as web servers and networked services.

#### 4. Standard Library

- Go's standard library is robust and extensive, covering a wide range of functionalities needed for backend development, including HTTP servers, file I/O, and cryptography.
- The standard library is well-documented and designed to be easy to use.

#### 5. Strong Community and Ecosystem

- Go has a growing and active community that contributes to a rich ecosystem of libraries and tools.
- There are many frameworks and libraries available for web development, database interaction, testing, and more.

#### 6. Cross-Platform Compatibility

- Go supports cross-compilation, allowing you to build binaries for multiple platforms from a single codebase.
- This makes deployment across different environments straightforward.

#### 7. Built-in Tooling

- Go comes with built-in tools for testing (go test), formatting (go fmt), and dependency management (go mod).
- These tools promote good coding practices and help maintain code quality.

### What Makes Go Different from Other Languages?

#### 1. Concurrency Model:

- Go's goroutines and channels offer a simpler and more efficient way to handle concurrency compared to traditional threading models in languages like Java or Python.

#### 2. Compilation Speed:

- Go compiles quickly, which speeds up the development cycle and allows for rapid iterations.

#### 3. Simplicity in Design:

- Go was designed to avoid complexity. It lacks many features found in other languages (like inheritance, generics, and assertions), which can be seen as a feature that reduces potential bugs and makes the language easier to learn and use.

#### 4. Garbage Collection:

- Go includes garbage collection, which simplifies memory management compared to languages like C or C++.

### What is Concurrency?

##### Explanation:

- Concurrency means executing multiple tasks simultaneously, but not necessarily at the same instant. It's about dealing with many things at once and making progress on several tasks concurrently.

##### Example 1: Simple Concurrency with Goroutines

- In Go, concurrency is achieved using goroutines, which are lightweight threads managed by the Go runtime.

```go
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(1 * time.Second) // Simulate a time-consuming task
    }
}

func main() {
    go printNumbers() // Start a new goroutine
    fmt.Println("Goroutine started")
    time.Sleep(6 * time.Second) // Give the goroutine time to finish
}
```

##### Example 2: Synchronizing Goroutines with Channels

- Channels provide a way for goroutines to communicate and synchronize their execution.

```go
package main

import (
    "fmt"
)

func sum(a, b int, result chan int) {
    result <- a + b // Send the result to the channel
}

func main() {
    result := make(chan int)
    go sum(1, 2, result)
    fmt.Println("Sum:", <-result) // Receive the result from the channel
}
```

###### In the examples:

- Goroutines: The go keyword starts a new goroutine.
- Channels: Channels are used to communicate between goroutines.

#### Key Concepts of Concurrency in Go

##### 1. Goroutines:

- Lightweight threads managed by the Go runtime.
- Syntax: go functionName()

##### 2. Channels:

- Used for communication between goroutines.
- Syntax: ch := make(chan int)

###### Example with Multiple Goroutines and Channels

```go
package main

import (
    "fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        results <- j * 2 // Simulate work by multiplying the job number
        fmt.Printf("Worker %d finished job %d\n", id, j)
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 worker goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs and close the channel
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results from workers
    for a := 1; a <= 5; a++ {
        fmt.Println("Result:", <-results)
    }
}
```

###### In the examples:

- Worker Function: Each worker receives jobs from the jobs channel and sends results to the results channel.
- Main Function: Starts worker goroutines, sends jobs to the jobs channel, and collects results from the results channel.

### What is go mod?

- go mod is a tool introduced in Go 1.11 to manage dependencies in Go projects. It helps you keep track of the libraries and packages your project depends on, ensuring consistent builds by specifying exact versions of these dependencies.

##### 1. Module Initialization:

- Initializes a new module in the current directory, creating a go.mod file.
- Command: `go mod init`

```go
go mod init example.com/myproject

```

##### 2. Add a dependency:

- Automatically downloads and installs the dependencies specified in your go.mod file.
  -Command: `go mod tidy`, `go get`

```go
go get github.com/gin-gonic/gin

```

- Note: We can compare `go mod` with node.js `package.json`

##### Comparing go mod with Node.js's package.json

| Feature             |         go mod         |          package.json |
| :------------------ | :--------------------: | --------------------: |
| Initialization      |      go mod init       |              npm init |
| Dependency File     |         go.mod         |          package.json |
| Adding Dependencies |    go get <package>    | npm install <package> |
| Module Caching      |    $GOPATH/pkg/mod     |          node_modules |
| Command to Install  | Implicit with go build |           npm install |

- go mod is similar to package.json in that both are used to manage dependencies in their respective ecosystems (Go and Node.js).
- go mod provides a file (go.mod) to list dependencies and their versions, similar to how package.json does for Node.js.
- Both tools help ensure consistent builds by keeping track of exact dependency versions and provide commands to add, update, and manage dependencies

### Variables:

- In Go, there are two ways to declare a variable

##### 1. Use the `var` keyword, followed by variable name and type

```go
package main
import ("fmt")

var name = "Jaddu" // Variable One

func main() {
 var fruit = "Apple" // Variable Two
 fmt.Println(name)
 fmt.Println(fruit)
}
```

##### 2. Use the := sign, followed by the variable value

```go
package main
import ("fmt")

func main() {
 varOne := 100
 varTwo := 2
 fmt.Println(varOne)
 fmt.Println(varTwo)
}
```

##### 2. Use the `const` keyword, followed by variable name and type

- When we use the `const` keyword can not change the value of the variable. It is meant to be constant

```go
package main
import ("fmt")

func main() {
 const countryName = "India"
 const age = 18
 fmt.Println(countryName)
 fmt.Println(age)
}
```

#### Assigning Types + Type Inferred

```go
package main
import ("fmt")

func main() {
 var fruit string = "Mango" // type is string
 var user = "Jaddu" // type is inferred
 number := 2 // type is inferred

 fmt.Println(fruit)
 fmt.Println(user)
 fmt.Println(number)
}
```

#### Go Multiple Variable Declaration

```go
package main
import ("fmt")

func main() {
 var one, two, three, four, five int = 1, 2, 3, 4, 5

 fmt.Println(one)
 fmt.Println(two)
 fmt.Println(three)
 fmt.Println(four)
}
```

### Go Variable Declaration in a Block

```go
package main
import ("fmt")

func main() {
  var (
   num int
   number int = 1
   greetings string = "hello"
  )

 fmt.Println(num)
 fmt.Println(number)
 fmt.Println(greetings)
}
```

#### Go Variable Naming Rules:

##### 1. A variable name must start with a letter or an underscore character (\_)

##### 2. A variable name cannot start with a digit

##### 3. A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and \_ )

##### 4. Variable names are case-sensitive (age, Age and AGE are three different variables)

##### 5. There is no limit on the length of the variable name

##### 6. A variable name cannot contain spaces

##### 7. The variable name cannot be any Go keywords I need this kind of explanation along with code example to learn go in readme syntax for each topic in go

##### Explanation:

- var: Keyword used to declare a variable.
- name: Variable name (follows naming conventions).
- string: Data type stored in the variable (name).
- =: Assignment operator, assigns the value "John Doe" to name.

### User Inputs in Go:

##### In Go, you can take user input using the `fmt` package. The `Scan` functions allow you to read input from standard input (stdin).

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    // Scanln reads a line from stdin
    // you can use Scan also to read
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}

```

##### Example: Reading Multiple Inputs

```go
package main

import "fmt"

func main() {
    var age int
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

```

### Pointers in Go:

- A pointer is a variable that points to the memory address of another variable.
- Pointers in Go allow you to store and manipulate the memory addresses of variables. This can be useful for several reasons, such as passing large structures efficiently, modifying the value of a variable inside a function, or working with data that needs to be shared across different parts of a program.
- In Go, the `*` symbol is used to declare a pointer. A pointer is a variable that holds the memory address of another variable. Pointers are useful because they allow functions to modify the value of a variable directly, rather than working with a copy of the variable.

#### Basic Concepts

##### 1. Pointer Declaration: A pointer holds the memory address of a variable.

##### 2. Address Operator (&): Used to get the memory address of a variable.

##### 3. Dereference Operator (\*): Used to access the value stored at the memory address.

#### Example: Basic Pointer Usage

##### 1. Declare a Pointer: Use \* to declare a pointer type.

##### 2. Get Address: Use & to get the address of a variable.

##### 3. Dereference: Use \* to access the value at the pointer's address.

```go
package main

import "fmt"

func main() {
    // Declare a variable
    var a int = 42

    // Declare a pointer that holds the address of variable a
    var p *int = &a

    // Print the value of a
    fmt.Println("Value of a:", a)

    // Print the address of a
    fmt.Println("Address of a:", &a)

    // Print the value of the pointer (which is the address of a)
    fmt.Println("Pointer p points to address:", p)

    // Dereference the pointer to get the value stored at the address
    fmt.Println("Value at the address p points to:", *p)
}

```

### Arrays in Go:

##### An array in Go is a collection of elements of the same type that are stored in a contiguous block of memory. Arrays are of fixed size, meaning once you define the size of an array, it cannot be changed.

##### Declaration without Initialization

- You can declare an array without initializing its elements. All elements will be set to the zero value for the type (e.g., 0 for integers, "" for strings).

```go
package main

import "fmt"

func main() {
    var arr [5]int
    fmt.Println(arr)
    // Output: [0 0 0 0 0]
}

```

##### Declaration with Initialization

- You can declare an array and initialize it at the same time.

```go
package main

import "fmt"

func main() {
    var arr = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)
    // Output: [1 2 3 4 5]
}


```

- You can also use the `:=` syntax for shorthand declaration and initialization

```go
package main

import "fmt"

func main() {
   arr := [5]int{1, 2, 3, 4, 5}
   fmt.Println(arr)
   // Output: [1 2 3 4 5]
}

```

##### Array Length

- You can get the length of an array using the `len` function.

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(len(arr))
    // Output: 5
}

```

##### Iterating Over Arrays

- You can use a `for` loop or a `for range` loop to iterate over the elements of an array.

##### Using for Loop

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}

```

##### Using for range Loop

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}

```

### Slices in Go:

##### Slices are a fundamental data structure in Go that provide a flexible and powerful way to work with collections of data. Unlike arrays, slices can grow and shrink in size (slices are dynamically sized), making them more versatile for most use cases.

##### Creating a Slice

##### 1. From an Array

- You can create a slice from an existing array by specifying a range.

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4] // Creates a slice from the second to the fourth element
    fmt.Println(slice) // Output: [2 3 4]
}

```

##### 2. Using the `make` Function

- The make function allows you to create a slice with a specified length and capacity. The length is the number of elements in the slice, and the capacity is the number of elements in the underlying array starting from the first element in the slice.

```go
package main

import "fmt"

func main() {
    slice := make([]int, 5, 10) // Length 5, Capacity 10
    fmt.Println(slice) // Output: [0 0 0 0 0]
    fmt.Println("Length:", len(slice)) // Output: Length: 5
    fmt.Println("Capacity:", cap(slice)) // Output: Capacity: 10
}

```

##### Slices Are References

- Slices are references to an underlying array. This means that if you modify the elements of a slice, the changes will be reflected in the underlying array and any other slices that refer to the same array.

```go
package main

import "fmt"

func main() {
   arr := [5]int{1, 2, 3, 4, 5}
   slice1 := arr[1:4]
   slice2 := arr[2:5]

   slice1[0] = 10
   fmt.Println(slice1) // Output: [10 3 4]
   fmt.Println(slice2) // Output: [3 4 5]
   fmt.Println(arr)    // Output: [1 10 3 4 5]
}

```

##### Appending to Slices

- The append function allows you to add elements to a slice. If the slice’s capacity is exceeded, a new underlying array is allocated.

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    slice = append(slice, 4, 5)
    fmt.Println(slice) // Output: [1 2 3 4 5]
}

```

##### Slicing Slices

- You can create a new slice from an existing slice by specifying a range.

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    newSlice := slice[1:4]
    fmt.Println(newSlice) // Output: [2 3 4]
}

```

##### Iterating Over Slices

- You can use a `for` loop or a `for range` loop to iterate over the elements of a slice.

##### Using for Loop

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }
}

```

##### Using for range Loop

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    for index, value := range slice {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

##### Summary:

- `Slices` in Go are more flexible and dynamic than arrays.
- Creating Slices: From arrays, using `make`, and slice literals.
- Slices Are `References`: Modifying one slice can affect others and the underlying array.
- `Appending`: Use append to add elements.
- Copying: Use `copy` to copy elements between slices.
- Slicing: Create new slices from existing ones.
- Iterating: Use `for` or `for range` loops.

### Maps in Go:

##### Maps in Go are a built-in data structure that associates keys with values. They provide a way to store and retrieve data efficiently, similar to dictionaries in Python or Objects in Js.

- A map is an unordered collection of key-value pairs. Each key in a map is unique, and you can use the key to look up the corresponding value.
- Default value of a map is nil

##### Creating and Initializing Maps

##### Using Map Literals

- You can create and initialize a map using a map literal.

```go
package main

import "fmt"

func main() {
    // Create and initialize a map using a map literal
    m := map[string]int{
        "Jaddu": 24,
        "Bob":   25,
    }
    fmt.Println(m) // Output: map[Alice:30 Bob:25]
}

```

##### Using the `make` Function

- You can create a map using the make function, specifying the key type and value type.

```go
package main

import "fmt"

func main() {
    // Create a map with string keys and int values
    m := make(map[string]int)
    fmt.Println(m) // Output: map[]
}

```

##### Adding and Retrieving Values

##### Adding Values

- You can add values to a map using the syntax `map[key] = value`.

```go
package main

import "fmt"

func main() {
   m := make(map[string]int)
   m["Jaddu"] = 24
   m["Bob"] = 25
   fmt.Println(m)
   // Output: map[Jaddu:24 Bob:25]
}


```

##### Retrieving Values

- You can retrieve values from a map using the key.

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "Jaddu": 24,
        "Bob":   25,
    }
    age := m["Jaddu"]
    fmt.Println("Jaddu's age:", age)
     // Output: Jaddu's age: 24
}

```

##### Checking if a Key Exists

- When retrieving a value from a map, you can check if the key exists by using the second return value from the map lookup.

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "Jaddu": 24,
        "Bob":   25,
    }
    age, ok := m["Charlie"]
    if ok {
        fmt.Println("Charlie's age:", age)
    } else {
        fmt.Println("Charlie is not in the map")
    }
}

```

##### Deleting a Key

- You can delete a key-value pair from a map using the `delete` function.

##### Using for Loop

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "Jaddu": 24,
        "Bob":   25,
    }
    delete(m, "Bob")
    fmt.Println(m)
    // Output: map[Jaddu:24]
}

```

##### Iterating Over a Map

- You can iterate over a map using a for range loop.

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }
}

```

##### Summary:

- `Maps` in Go are used to store key-value pairs.
- Creating Maps: Use `make` or `map literals`.
- Adding/Retrieving Values: Use `map[key] = value` and `map[key]`.
- Checking Existence: Use the second return value from a map lookup.
- Deleting Keys: Use the `delete` function.
- Iterating: Use `for rang`e loops.

### Structs in Go:

##### Structs are one of the most important and commonly used data structures in Go. They allow you to group together related data, which can be of different types, into a single unit. You can think of a struct as a way to create your own custom data types.

- A struct is a collection of fields. Each field has a name and a type. Structs are used to group related data together.
- You can compare `Struct` with `Schema` in Mongoose

##### Defining a Struct

- To define a struct, use the `type` and `struct` keywords.

```go
package main

import "fmt"

// Define a struct type named Person
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create a new instance of the Person struct
    var p1 Person
    p1.Name = "Jaddu"
    p1.Age = 24
    p1.City = "Bangalore"

    fmt.Println(p1)
    // Output: {Jaddu 24 Bangalore}
}

```

##### Creating and Initializing Structs

##### Using `Struct` Literals

- You can create and initialize a struct using a `struct` literal.

```go
package main

import "fmt"

// Define the Person struct
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Initialize a struct using a struct literal
    p2 := Person{Name: "Sara", Age: 22, City: "Los Angeles"}
    fmt.Println(p2) // Output: {Sara 25 Los Angeles}
}

```

##### Using `new` Function

- You can also create a pointer to a struct using the `new` function.

```go
package main

import "fmt"

// Define the Person struct
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Create a pointer to a new Person struct
    p3 := new(Person)
    p3.Name = "Charlie"
    p3.Age = 35
    p3.City = "New York"

    fmt.Println(*p3) // Output: {Charlie 35 New York"}
}


```

##### Accessing and Modifying Struct Fields

- You can access and modify the fields of a struct using the `dot notation`.

```go
package main

import "fmt"

// Define the Person struct
type Person struct {
   Name string
   Age  int
   City string
}

func main() {
   // Create and initialize a Person struct
   p := Person{Name: "Dave", Age: 40, City: "San Francisco"}

   // Access fields
   fmt.Println("Name:", p.Name)
   // Output: Name: Dave
   fmt.Println("Age:", p.Age)
   // Output: Age: 40
   fmt.Println("City:", p.City)
   // Output: City: San Francisco

   // Modify fields
   p.Age = 41
   fmt.Println("Updated Age:", p.Age)
   // Output: Updated Age: 41
}

```

##### Nested Structs

- Structs can be nested, meaning you can have a struct as a field within another struct.

```go
package main

import "fmt"

// Define the Address struct
type Address struct {
    Street string
    City   string
    Zip    string
}

// Define the Person struct with an embedded Address struct
type Person struct {
    Name    string
    Age     int
    Address Address
}

func main() {
    // Create and initialize a nested struct
    p := Person{
        Name: "Eve",
        Age:  28,
        Address: Address{
            Street: "123 Main St",
            City:   "Boston",
            Zip:    "02118",
        },
    }

    fmt.Println(p)
    // Output: {Eve 28 {123 Main St Boston 02118}}

    // Access nested fields
    fmt.Println("City:", p.Address.City)
    // Output: City: Boston
}

```

##### Capitalization Rules

- In Go, the visibility of struct fields and methods is determined by their capitalization. By following this convention, you control whether a field or method is accessible outside of the package where it is defined.

##### 1.Exported (Public) Fields and Methods:

- Fields and methods that start with an uppercase letter are exported, meaning they are accessible from other packages.

##### 2.Unexported (Private) Fields and Methods:

- Fields and methods that start with a lowercase letter are unexported, meaning they are only accessible within the package where they are defined.

```go

package main

import (
	"fmt"
)

type Person struct {
	Name    string // Exported field (public)
	age     int    // Unexported field (private)
}

// Exported method
func (p Person) GetAge() int {
	return p.age
}

// Unexported method
func (p *Person) setAge(age int) {
	p.age = age
}

func main() {
	p := Person{Name: "Alice", age: 30}

	// Accessing exported field
	fmt.Println("Name:", p.Name)

	// Accessing unexported field directly (will cause a compilation error)
	// fmt.Println("Age:", p.age)

	// Accessing unexported field via exported method
	fmt.Println("Age:", p.GetAge())

	// Setting age using an unexported method within the same package
	p.setAge(35)
	fmt.Println("New Age:", p.GetAge())
}

```

- Use uppercase for fields and methods that need to be accessed from outside the package.
- Use lowercase for fields and methods that should be hidden within the package.

##### Methods on Structs

- You can define methods on structs to add behavior to them. A method is a function with a special receiver argument.

```go
package main

import "fmt"

// Define the Person struct
type Person struct {
    Name string
    Age  int
}

// Define a method on the Person struct
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create and initialize a Person struct
    p := Person{Name: "Jaddu", Age: 20}

    // Call the Greet method
    p.Greet()
    // Output: Hello, my name is Jaddu and I am 20 years old.
}

```

##### Struct Tags:

- Struct tags in Go are metadata attached to struct fields. They provide instructions to the Go runtime about how to handle the fields when encoding (marshaling) or decoding (unmarshaling) structs to/from different formats like JSON, XML, or database rows.

##### Example Struct with Tags

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"`
    Password string `json:"-"`
}

```

##### Explanation of Struct Tags

##### 1. Basic Syntax: Struct tags are written after the field type and enclosed in backticks `json:"tag"`.

##### 2. Key-Value Pairs: Struct tags consist of key-value pairs separated by colons` :`. Each key-value pair provides a specific instruction.

- `json:"id"`: This instructs the JSON encoder to use `"id"` as the key when marshaling this field into JSON. For example, `{ID: 1}` becomes `{"id": 1}` in JSON.
- `json:"name"`: Similarly, this specifies that "`name"` should be used as the key in JSON.
- `json:"email,omitempty"`: The omitempty option tells the JSON encoder to omit this field from the JSON output if it's empty (`""`). Useful for optional fields.
- `json:"-"`: The dash`-` indicates that this field should be ignored by the JSON encoder. It won't appear in the JSON output.

##### Why Use Struct Tags?

- Serialization: Struct tags are crucial for encoding Go structs into different formats like JSON, XML, etc., and vice versa.
- Flexibility: They allow fine-grained control over how struct fields are represented in different data formats, ensuring compatibility and efficiency.

##### Summary:

- Structs: Collections of fields that group related data.
- Defining Structs: Use the `type` and `struct` keywords.
- Creating Structs: Use struct literals or the `new` function.
- Accessing/Modifying Fields: Use the dot notation.
- Nested Structs: Structs can contain other structs.
- Methods on Structs: Define methods to add behavior to structs

### Functions in Go:

##### In Go, a function is a block of code that can be executed multiple times from different parts of your program. It's a way to group a set of statements together to perform a specific task.

- To define a function in Go, you use the `func` keyword, followed by the function name, parameters (if any), and the return type (if any).

##### Example of a function in Go:

```go
package main

import "fmt"

// Define a function named greet that takes a string parameter and returns a string
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}

func main() {
    // Call the greet function and print the result
    greet("John")
    greet("Jane")
   // Output: Hello, John!
           // Hello, Jane!
}

```

##### Function Declaration

- A function declaration consists of the `func` keyword, followed by the function name, parameters in parentheses, and the function body in curly braces.

```go
func functionName(param1 type1, param2 type2, ...) returnType {
    // function body
}
```

##### Function Parameters

- Functions can take zero or more parameters, which are specified in the parentheses after the function name. Parameters are separated by commas.

```go
func add(x int, y int) int {
   return x + y
}

```

- In this example, the add function takes two int parameters, x and y.

##### Function Return Types

- A function can return one or more values, which are specified after the parameter list. If a function returns no values, the return type is omitted.

```go
func add(x int, y int) int {
    return x + y
}

func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}

```

- In the first example, the add function returns an int value. In the second example, the greet function returns no values.

##### Multiple Return Values:

- Go functions can return multiple values, which is useful for returning both a result and an error status, for example.

```go
func divide(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("division by zero")
    }
    return dividend / divisor, nil
}

```

##### Function Calls

- To call a function, simply use the function name followed by parentheses containing the arguments.

```go
result := add(2, 3)
fmt.Println(result)
// Output: 5

greet("John")

```

##### Function Scope

- Variables declared inside a function are only accessible within that function. This is known as the function's scope.

```go
func main() {
    x := 10
    fmt.Println(x)
    // Output: 10

    func inner() {
        y := 20
        fmt.Println(y)
        // Output: 20
    }

    fmt.Println(y)
     // Error: y is not defined
}
```

- In this example, the variable x is accessible in the main function, but the variable y is only accessible within the inner function.

##### Closures

- A closure is a function that has access to its own scope and the scope of its surrounding functions.

```go
func outer() func() int {
    x := 10
    return func() int {
        x++
        return x
    }
}

func main() {
    inner := outer()
    fmt.Println(inner()) // Output: 11
    fmt.Println(inner()) // Output: 12
}

```

- In this example, the outer function returns a closure that has access to the x variable. The closure is called multiple times, and each time it increments the x variable and returns its value

##### Anonymous Functions

- Functions in Go can be defined without a name and can be assigned to variables or passed as arguments.

```go
package main

import "fmt"

func main() {
    // Define and call an anonymous function
    func() {
        fmt.Println("Hello, anonymous function!")
    }()

    // Define an anonymous function and assign it to a variable
    add := func(a int, b int) int {
        return a + b
    }

    // Call the anonymous function
    fmt.Println("Sum:", add(3, 4)) // Output: Sum: 7
}


```

##### Higher-Order Functions

- A higher-order function is a function that takes another function as an argument or returns a function as a result.

```go
func apply(f func(int) int, x int) int {
    return f(x)
}

func double(x int) int {
    return x * 2
}

func main() {
    result := apply(double, 5)
    fmt.Println(result) // Output: 10
}


```

- In this example, the apply function takes a function `f` and an integer `x` as arguments. It calls the function `f` with the argument x and returns the result. The `double` function is passed as an argument to the `apply` function, which doubles the input value.

##### Panic Function in Go:

- The `panic` function in Go is used to cause a program to terminate immediately and print an error message. It is generally used to handle unrecoverable errors that should stop the normal execution of the program. When `panic` is called, the program stops executing the current function, runs any deferred functions, and then terminates.

##### Example:

```go
package main

import "fmt"

func main() {
   fmt.Println("Start of main function")

   panic("Something went wrong!")

   fmt.Println("This line will never be executed")
}
// O/P:
Start of main function
panic: Something went wrong!

goroutine 1 [running]:
main.main()
/path/to/your/file.go:8 +0x40
exit status 2

```

##### 1. Usage

- `panic` should be used for critical errors where the program cannot continue.
- It is typically used in situations where something has gone seriously wrong, such as encountering an unexpected condition that makes further execution impossible or dangerous.

##### 2. Behavior:

- When `panic` is called, the program stops executing the current function and begins unwinding the stack.
- Deferred functions (functions scheduled to run with defer) are executed in the reverse order they were deferred.
- After running all deferred functions, the program exits and prints the panic message along with a stack trace to the console.

##### Example with `defer`

```go
package main

import "fmt"

func main() {
   defer fmt.Println("This will be printed last, before the program exits")

   fmt.Println("Start of main function")

   panic("Something went wrong!")

   fmt.Println("This line will never be executed")
}

```

### Methods in Go:

- In Go, methods are functions with a special receiver argument. The receiver can be a value or a pointer to a struct. Methods allow you to define behaviors associated with your data types.

##### Method Syntax:

```go
func (receiverType Receiver) methodName(parameters) returnType {
   // method body
}

```

##### Example:

```go
package main

import (
	"fmt"
)

// Defining the Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Method to get the full name of the person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Method to check if the person is an adult
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Method to increment the age of the person
// Notice the pointer receiver to modify the original struct
func (p *Person) IncrementAge() {
	p.Age++
}

func main() {
	// Creating an instance of Person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       20,
	}

	// Calling the FullName method
	fmt.Println("Full Name:", person.FullName())

	// Calling the IsAdult method
	fmt.Println("Is Adult:", person.IsAdult())

	// Incrementing the age using the IncrementAge method
	person.IncrementAge()
	fmt.Println("New Age:", person.Age)
}


```

##### Real-Time Usage in API Development:

- In a real-world scenario, you might define methods on structs that represent entities in your application. For example, in an API for user management, you might have a User struct with methods for various operations.

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// Method to check if the email is valid (dummy example)
func (u User) IsValidEmail() bool {
	// Simplified email validation
	return len(u.Email) > 5 && u.Email[len(u.Email)-4:] == ".com"
}

// Method to hide the password from JSON response
func (u *User) HidePassword() {
	u.Password = ""
}

// Handler to get user information
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret",
	}

	// Hide the password before sending the response
	user.HidePassword()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUserHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

```

- By using methods, you can encapsulate behavior related to your data types and make your code more organized and modular.

### What is Defer:

- The `defer` statement in Go is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. Deferred function calls are executed in the reverse order that they were deferred, right before the surrounding function returns, regardless of whether the function returns normally or via a panic.
- Deferred functions are executed in Last In, First Out (LIFO) order
- This means that if you defer multiple function calls, the last one deferred will be the first to execute when the surrounding function returns.

##### Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Start of main function")

    defer fmt.Println("This is deferred")

    fmt.Println("End of main function")
}
O/P:
 Start of main function
 End of main function
 This is deferred


```

##### Multiple Deferred Calls:

```go
package main

import "fmt"

func main() {
    fmt.Println("Start of main function")

    defer fmt.Println("First deferred call")
    defer fmt.Println("Second deferred call")
    defer fmt.Println("Third deferred call")

    fmt.Println("End of main function")
}
O/P:
  Start of main function
  End of main function
  Third deferred call
  Second deferred call
  First deferred call

```

### Advanced:

### Interfaces in GO:

- In Go, an interface is a collection of methods that a type can implement. It's a way to define a contract that a type must adhere to, without caring about the underlying implementation.
- Interfaces are a way to define a set of method signatures that a type must implement. Unlike other languages, Go’s interfaces are implemented implicitly. This means that a type does not need to explicitly declare that it implements an interface, it just needs to provide the methods that the interface requires

##### Syntax

- An interface in Go specifies a method set, which is a collection of method signatures. Any type that has methods matching the method set of an interface is considered to implement that interface.

```go

type InterfaceName interface {
    MethodName1(param1 type1, param2 type2) returnType1
    MethodName2(param1 type1) returnType2
}

```

##### Example

```go
package main

import "fmt"

// Define a simple interface named 'Describer'
type Describer interface {
    Describe() string
}

// Implement the interface for a 'Person'
type Person struct {
    Name string
    Age  string
}

// Implement the 'Describe' method for 'Person'
func (p Person) Describe() string {
    return "Hi, I'm " + p.Name + "and my age is " + p.Age
}

// Implement the interface for a 'Car'
type Car struct {
    Model string
}

// Implement the 'Describe' method for 'Car'
func (c Car) Describe() string {
    return "This is a " + c.Model + " car"
}

// Function to display description using the interface
func displayDescription(d Describer) {
    fmt.Println(d.Describe())
}

func main() {
    // Create instances of Person and Car
    person := Person{Name: "Jaddu"}
    car := Car{Model: "BMW"}

    // Use the Describe method through the Describer interface
    displayDescription(person)
    displayDescription(car)
}
```

##### Explanation:

- We define an interface `Describer` with a single method `Describe`.
- We define a struct type `Person` and `Car`.
- We implement the `Describe` method for the `Person` type and `Car` type, which makes `Person` & `Car` satisfy the `Describer` interface.
- Create instances of `Person` and `Car`.
- Use the `Describe` method through the `Describer`

### Backend:

### JSON in GO:

- JSON (JavaScript Object Notation) is a lightweight data format used for exchanging data between systems. It is human-readable, easy to parse, and widely used in APIs.

```go

{
  "name": "Jabeed",
  "age": 25,
  "skills": ["Go", "Gin", "MongoDB"],
  "isEmployed": true
}

```

- Strings → "name": "Jabeed"
- Numbers → "age": 25
- Boolean → "isEmployed": true
- Arrays → "skills": ["Go", "Gin", "MongoDB"]
- Nested Objects → { "address": { "city": "Hyderabad", "country": "India" } }

##### Why is JSON Needed?

- Used in REST APIs to send and receive data.
- Platform-independent (works with Go, Java, Python, JavaScript, etc.).
- Replaces XML, which is bulky and harder to parse.
- Easily readable and structured.

#### 1. Converting Go Structs to JSON (Marshaling)

- In Go, we convert Go structs to JSON using json.Marshal(). This is called Marshaling.

##### Example: Convert Go Struct to JSON

```go

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Skills     []string `json:"skills"`
	IsEmployed bool     `json:"isEmployed"`
}

func main() {
	// Create an instance of User struct
	user := User{
		Name:       "Jabeed",
		Age:        25,
		Skills:     []string{"Go", "Gin", "MongoDB"},
		IsEmployed: true,
	}

	// Convert Go struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print JSON as string
	fmt.Println(string(jsonData))
}


```

##### output

```go
{"name":"Jabeed","age":25,"skills":["Go","Gin","MongoDB"],"isEmployed":true}

```

#### 2. Formatting JSON Output (Pretty Print)

- By default, json.Marshal() produces a compact JSON.
- For pretty printing, we use json.MarshalIndent().

##### Example: Formatting JSON Output

```go

jsonPretty, _ := json.MarshalIndent(user, "", "  ")
fmt.Println(string(jsonPretty))

```

##### output

```go
{
  "name": "Jabeed",
  "age": 25,
  "skills": [
    "Go",
    "Gin",
    "MongoDB"
  ],
  "isEmployed": true
}


```

#### 3. Converting JSON to Go Struct (Unmarshaling)

- When an API sends JSON data, we convert JSON into Go structs using json.Unmarshal(). This is called Unmarshaling.

##### Example: Convert JSON to Go Struct

```go

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Skills     []string `json:"skills"`
	IsEmployed bool     `json:"isEmployed"`
}

func main() {
	// Sample JSON (coming from an API)
	jsonString := `{"name":"Jabeed","age":25,"skills":["Go","Gin","MongoDB"],"isEmployed":true}`

	// Convert JSON string to Go struct
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print converted struct
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Skills:", user.Skills)
	fmt.Println("Is Employed:", user.IsEmployed)
}


```

##### output

```go
Name: Jabeed
Age: 25
Skills: [Go Gin MongoDB]
Is Employed: true

```

#### 4. Handling Nested JSON Objects

- By default, json.Marshal() produces a compact JSON.
- For pretty printing, we use json.MarshalIndent().

##### Example: Nested JSON to Go Struct

```go

package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonString := `{"name":"Jabeed","age":25,"address":{"city":"Hyderabad","country":"India"}}`

	var user User
	json.Unmarshal([]byte(jsonString), &user)

	fmt.Println("Name:", user.Name)
	fmt.Println("City:", user.Address.City)
	fmt.Println("Country:", user.Address.Country)
}


```

##### output

```go
Name: Jabeed
City: Hyderabad
Country: India



```

#### 5. Handling JSON with Dynamic Keys (Map)

- If JSON has unknown keys, we use map[string]interface{}

##### Example: Parsing JSON into a Map

```go

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{"id": 101, "name": "Jabeed", "active": true, "score": 99.5}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonString), &data)

	fmt.Println("ID:", data["id"])
	fmt.Println("Name:", data["name"])
	fmt.Println("Active:", data["active"])
	fmt.Println("Score:", data["score"])
}

```

##### output

```go
ID: 101
Name: Jabeed
Active: true
Score: 99.5

```

#### 6. Ignoring Fields & Custom Naming

- Ignore a field by using json:"-"
- Change JSON key name using struct tags

##### Example: Custom JSON Struct Tags

```go

type User struct {
	Name       string `json:"full_name"`
	Age        int    `json:"-"`
	Occupation string `json:"job,omitempty"`
}

```

- 1️⃣ "full_name" → Changes "Name" key to "full_name" in JSON.
- 2️⃣ "-" → Skips "Age" in JSON output.
- 3️⃣ "omitempty" → Skips "Occupation" if empty.

#### 7. Encoding JSON Response in an API

- In Gin, we use c.JSON() to send JSON responses.

##### Example:

```go

router.GET("/user", func(c *gin.Context) {
	user := User{Name: "Jabeed", Age: 25, Skills: []string{"Go", "Gin"}}
	c.JSON(http.StatusOK, user)
})



```

##### output

```go
{
  "name": "Jabeed",
  "age": 25,
  "skills": ["Go", "Gin"]
}


```

##### Explanation:

- We define an interface `Describer` with a single method `Describe`.
- We define a struct type `Person` and `Car`.
- We implement the `Describe` method for the `Person` type and `Car` type, which makes `Person` & `Car` satisfy the `Describer` interface.
- Create instances of `Person` and `Car`.
- Use the `Describe` method through the `Describer` interface
