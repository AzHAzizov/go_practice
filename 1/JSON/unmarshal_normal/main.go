package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User
}

type User struct {
	Name   string `json:"name"`
	Role   string `json:"type"`
	Age    int    `json:"age"`
	Social Social
}

type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

func main() {

	jsonFile, err := os.Open("users.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	var users interface{}

	fmt.Println("Application for work with json works CORRECT")

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)

	fmt.Println(users)

	// for _, user := range users.Users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println("________________________")
	// }
}
