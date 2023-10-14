package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id     int    `json:"id"`
	Date   string `json:"create_time"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Author int    `json:"author"`
}

func main() {
	args := []byte(`{"name":"name", "create_time":"2023-10-14","author":14,"id":1}`)

	var user User
	err := json.Unmarshal(args, &user)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	fmt.Println(user.Id)
	fmt.Println(user.Name)
}
