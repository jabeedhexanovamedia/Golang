### Backend:

### JSON in GO:

- JSON (JavaScript Object Notation) is a lightweight data format used for exchanging data between systems. It is human-readable, easy to parse, and widely used in APIs.
- JSON (JavaScript Object Notation) is a fundamental part of modern web development, especially in APIs. It’s lightweight, human-readable, and easy for machines to parse and generate.
- Go provides the encoding/json package to work with JSON.

##### What is JSON?

- JSON is a text-based data format used to represent structured data.
- It’s commonly used to transmit data between a server and a web application (e.g., in REST APIs).
- JSON is language-independent, meaning it can be used with almost any programming language.

##### Why Do We Need JSON?

- Data Exchange: JSON is the de facto standard for APIs. It allows clients and servers to communicate efficiently.
- Human-Readable: Unlike binary formats, JSON is easy to read and debug.
- Lightweight: JSON has a small footprint compared to XML.
- Wide Support: Almost every programming language has built-in support for JSON.

##### JSON Syntax

- JSON is made up of key-value pairs and supports the following data types:
- Strings: "name": "John"
- Numbers: "age": 30
- Booleans: "isStudent": true
- Arrays: "hobbies": ["reading", "coding"]
- Objects: "address": {"city": "New York", "zip": "10001"}

##### Examples:

```go

{
  "name": "John",
  "age": 30,
  "isStudent": false,
  "hobbies": ["reading", "coding"],
  "address": {
    "city": "New York",
    "zip": "10001"
  }
}

```

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
- Sometimes, you may not know the structure of the JSON in advance. In such cases, you can use a map[string]interface{}.

##### Example: Parsing JSON into a Map

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
		"name": "John",
		"age": 30,
		"address": {
			"city": "New York",
			"zip": "10001"
		},
		"hobbies": ["reading", "coding"]
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Address:", data["address"].(map[string]interface{}))
	fmt.Println("Hobbies:", data["hobbies"].([]interface{}))
}
```

```go
Name: John
Age: 30
Address: map[city:New York zip:10001]
Hobbies: [reading coding]
```

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

##### What is map[string]interface{}?

- In Go, interface{} is a special type that can hold any value (similar to Object in Java or any in TypeScript). When you use map[string]interface{}, you’re creating a map where:

- Keys are strings (e.g., "name", "age").

- Values can be of any type (e.g., string, int, bool, map[string]interface{}, etc.).

- This is particularly useful when working with JSON data where the structure is dynamic or unknown.

##### Why Use map[string]interface{} for JSON?

- When you receive JSON data, you often know the structure in advance. For example:

```go
{
  "name": "John",
  "age": 30
}
```

- You can define a struct like this:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

- But what if the JSON structure is dynamic or unknown? For example:

```go
{
  "name": "John",
  "age": 30,
  "address": {
    "city": "New York",
    "zip": "10001"
  },
  "hobbies": ["reading", "coding"]
}
```

- Here, the JSON contains nested objects and arrays. If you don’t know the structure in advance, you can’t define a fixed struct. This is where map[string]interface{} comes in handy.

##### How to Use map[string]interface{} for JSON

##### Example: Unmarshalling Dynamic JSON

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
		"name": "John",
		"age": 30,
		"address": {
			"city": "New York",
			"zip": "10001"
		},
		"hobbies": ["reading", "coding"]
	}`

	// Declare a map to hold the JSON data
	var data map[string]interface{}

	// Unmarshal the JSON into the map
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Access the values
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])

	// Access nested objects
	address := data["address"].(map[string]interface{})
	fmt.Println("City:", address["city"])
	fmt.Println("Zip:", address["zip"])

	// Access arrays
	hobbies := data["hobbies"].([]interface{})
	fmt.Println("Hobbies:")
	for _, hobby := range hobbies {
		fmt.Println("-", hobby)
	}
}

```

```go
Name: John
Age: 30
City: New York
Zip: 10001
Hobbies:
- reading
- coding

```

##### How It Works:

##### 1. Unmarshalling:

- json.Unmarshal() reads the JSON data and populates the map[string]interface{}.

- Nested objects become map[string]interface{}.

- Arrays become []interface{}.

##### 2. Type Assertion:

- To access nested objects or arrays, you need to perform type assertion:

```go
address := data["address"].(map[string]interface{})
hobbies := data["hobbies"].([]interface{})
This tells Go that data["address"] is a map[string]interface{} and data["hobbies"] is a []interface{}.
```

##### 3. Accessing Values:

- Use keys to access values in the map.

- For nested structures, use type assertion to "unwrap" the data.

##### Use Cases for map[string]interface{}

1.Dynamic JSON:

- When the structure of the JSON is not known in advance.

- Example: APIs that return flexible or user-defined data.

2. Generic Data Handling:

- When you need to process data without knowing its exact structure.

- Example: Parsing configuration files or API responses.

3. Prototyping:

- When you’re quickly prototyping and don’t want to define structs.

4. Third-Party APIs:

- When working with APIs that return unpredictable or changing data structures.

##### Real-World Example: Handling API Responses

- Imagine you’re working with a weather API that returns dynamic data:

```go
{
  "location": "New York",
  "temperature": 22.5,
  "conditions": "Sunny",
  "forecast": {
    "tomorrow": {
      "high": 25,
      "low": 18
    }
  }
}
```

- You can handle this using map[string]interface{}:

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
		"location": "New York",
		"temperature": 22.5,
		"conditions": "Sunny",
		"forecast": {
			"tomorrow": {
				"high": 25,
				"low": 18
			}
		}
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Location:", data["location"])
	fmt.Println("Temperature:", data["temperature"])

	forecast := data["forecast"].(map[string]interface{})
	tomorrow := forecast["tomorrow"].(map[string]interface{})
	fmt.Println("Tomorrow's High:", tomorrow["high"])
	fmt.Println("Tomorrow's Low:", tomorrow["low"])
}
```

```go
Location: New York
Temperature: 22.5
Tomorrow's High: 25
Tomorrow's Low: 18
```

##### Key Takeaways

- Maps are key-value pairs in Go.

- map[string]interface{} is used for dynamic or unknown JSON structures.

- Type assertion is required to access nested data.

- Use map[string]interface{} when you need flexibility, but be aware of its limitations.

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
