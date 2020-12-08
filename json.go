package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type result struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Email struct {
		Personal string `json:"personal"`
		Work     string `json:"work"`
	}
	Phone string `json:"phone"`
}

func main() {
	someJSON := `{
		"name": "vireak",
		"id": 1,
		"email": {
			"personal": "someemail@gmail.com",
			"work": "work@somecompany.com"
		},
		"phone": "+85512345678"
	}`
	// unmarshal to a struct in case we know the json schema
	var res result
	if err := json.Unmarshal([]byte(someJSON), &res); err != nil {
		log.Fatal("Invalid JSON: ", err)
	}
	fmt.Println(res)
	fmt.Println(res.Email.Personal)

	// unmarshal to a map[string]interface{} in case we don't know the json schema
	data := make(map[string]interface{})
	if err := json.Unmarshal([]byte(someJSON), &data); err != nil {
		log.Fatal("Invalid JSON: ", err)
	}
	fmt.Println(data)
	fmt.Println(data["name"])
	fmt.Println(data["id"])
	fmt.Println(data["email"])

	// assert that the nested json data is of type (map[string]interface{}) https://golang.org/ref/spec#Type_assertions
	if _, ok := data["email"].(map[string]interface{}); !ok {
		log.Fatal("Invalid data schema!")
	}
	fmt.Println(data["email"].(map[string]interface{}))
	fmt.Println(data["email"].(map[string]interface{})["personal"])
	fmt.Println(data["phone"])

}
