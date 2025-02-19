package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skills []string `json:"skills"`
	IsEmployed bool `json:"is_employed"`
}

func main() {

	fmt.Println("JSON")

	// Create an instance of User struct

	user := User{
		Name:"Jaddu", 
		Age:24, 
		Skills:[]string{"Go", "AWS", "ReactJs"}, IsEmployed: true, 
	}

	fmt.Println("User instance created", user)

	// 1. Marshalling (Encoding)
	// Convert Go struct to JSON

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	fmt.Println("JSON converted data without format:", string(jsonData))

	// 2. Formatting JSON Output (Pretty Print)

	jsonPretty,_ := json.MarshalIndent(user, "", "")

	fmt.Println("JSON data Pretty Print:", string(jsonPretty))

	// 3. Converting JSON to Go Struct (Unmarshaling)

	// Sample JSON (coming from an API)

	jsonString := `{"name":"Jabeed","age":24,"skills":["Go","Gin","MongoDB"],"isEmployed":true}`

	// Convert JSON string to Go struct
	var user2 User

	err2 := json.Unmarshal([]byte(jsonString), &user2)
	if err2 != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// Print converted struct
	fmt.Println("Converted struct:", user2)

		


}