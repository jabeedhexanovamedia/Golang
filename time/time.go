package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from Time package")

	currentTime := time.Now()

	fmt.Println("current time without format: ", currentTime)

	// formattedTime := currentTime.Format("02-01-2006, Monday")
	formattedTime := currentTime.Format("02-01-2006, 15-04-05")

	// Time is represent in go as 02-01-2006 15-04-05 (DD-MM-YYYY HH:MM:SS)
	// Monday to get the current day name

	fmt.Println("Formatted time: ", formattedTime)
}