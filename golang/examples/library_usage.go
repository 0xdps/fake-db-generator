package main

import (
	"encoding/json"
	"fmt"

	"github.com/0xdps/fake-stack/golang/pkg/generator"
)

func main() {
	gen := generator.New()

	userFields := []generator.Field{
		{Name: "id", Generator: "uuid"},
		{Name: "username", Generator: "username"},
		{Name: "email", Generator: "email"},
		{Name: "first_name", Generator: "first_name"},
		{Name: "last_name", Generator: "last_name"},
		{Name: "avatar", Generator: "url"},
		{Name: "created_at", Generator: "past_date"},
	}

	users, err := gen.GenerateRecords(userFields, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	jsonData, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println("Generated Users:")
	fmt.Println(string(jsonData))
}
