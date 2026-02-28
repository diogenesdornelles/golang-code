package serializacao

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	u := User{1, "Alice"}

	// Struct -> JSON (Marshal)
	jsonData, _ := json.Marshal(u)
	fmt.Println(string(jsonData)) // {"id":1,"nome":"Alice"}

	// JSON -> Struct (Unmarshal)
	var u2 User
	json.Unmarshal(jsonData, &u2)
	fmt.Println(u2) // {1 Alice}

}
