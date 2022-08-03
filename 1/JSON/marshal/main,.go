package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  byte   `json:"sex"`
	Home Home
}

type Home struct {
	City       string `json:"city"`
	Street     string `json:"street"`
	HomeNumber string `json:"home"`
}

func main() {

	user1 := Human{
		Name: "User1",
		Age:  31,
		Sex:  'M',
		Home: Home{
			City:       "Alabama",
			Street:     "32 avenue",
			HomeNumber: "21a",
		},
	}

	bytedJson, err := json.Marshal(user1)

	if err != nil {
		log.Fatal(err)
	}

	var jsonData string = string(bytedJson)
	fmt.Println(jsonData, "Get user data as json")

	ioutil.WriteFile("user_data.json", bytedJson, 0664)

}
