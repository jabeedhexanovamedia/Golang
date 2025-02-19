package main

import "fmt"

func deferExample() {
    fmt.Println("Start of the program")

    // Use defer to schedule this line to run at the end of the function
    defer fmt.Println("This is the end of the program")

    fmt.Println("This is the middle of the program")
}