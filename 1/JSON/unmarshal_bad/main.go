package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
}
