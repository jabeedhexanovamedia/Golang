// package main

// import "fmt"

// func main() {
//     fmt.Println("Start of main function")

//     defer fmt.Println("This is deferred")

//     fmt.Println("End of main function")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
//     file, err := os.Open("example.txt")
//     if err != nil {
//         fmt.Println("Error opening file:", err)
//         return
//     }
//     defer file.Close()

//     // Read from file or perform other operations
//     fmt.Println("File opened successfully")
// }

package main

func main() {
    deferExample()
}

// to run: go run defer.go defer2.go