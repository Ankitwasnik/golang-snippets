package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	println("This file has most used json methods")

	user := User{Name: "ABC", Age: 23, Country: "USA"}

	userString, _ := StructToString(user)
	fmt.Printf("Marshalled user string %s\n", userString)

	userUnmarshalled, _ := StringToStruct(userString)
	fmt.Printf("Unmarshalled user %#v\n", userUnmarshalled)

}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

// Note: This function uses json tag present in the struct fields.
// So, fields which do not contain json tag will not be unmarshalled.
func StringToStruct(jsonString string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Note: This function uses json tag present in the struct fields.
// So, fields which do not contain json tag will not be marshalled.
func StructToString(user User) (string, error) {
	marshalledBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(marshalledBytes), nil
}
